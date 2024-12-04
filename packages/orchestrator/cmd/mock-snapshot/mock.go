package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"golang.org/x/sync/errgroup"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/dns"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/network"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/template"
	"github.com/e2b-dev/infra/packages/shared/pkg/grpc/orchestrator"
	"github.com/e2b-dev/infra/packages/shared/pkg/logs"
	"github.com/e2b-dev/infra/packages/shared/pkg/storage"
)

func main() {
	templateId := flag.String("template", "", "template id")
	buildId := flag.String("build", "", "build id")
	sandboxId := flag.String("sandbox", "", "sandbox id")
	keepAlive := flag.Int("alive", 0, "keep alive")
	count := flag.Int("count", 1, "number of serially spawned sandboxes")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() {
		<-done

		cancel()
	}()

	dnsServer := dns.New()
	go func() {
		log.Printf("Starting DNS server")

		err := dnsServer.Start("127.0.0.4:53")
		if err != nil {
			log.Fatalf("Failed running DNS server: %s\n", err.Error())
		}
	}()

	templateCache := template.NewCache(ctx)

	networkPool, err := network.NewPool(ctx, *count, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create network pool: %v\n", err)

		return
	}
	defer networkPool.Close()

	eg, ctx := errgroup.WithContext(ctx)

	for i := 0; i < *count; i++ {
		fmt.Println("--------------------------------")
		fmt.Printf("Starting sandbox %d\n", i)

		v := i

		err = mockSnapshot(
			ctx,
			*templateId,
			*buildId,
			*sandboxId+"-"+strconv.Itoa(v),
			dnsServer,
			time.Duration(*keepAlive)*time.Second,
			networkPool,
			templateCache,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to start sandbox: %v\n", err)
			return
		}
	}

	err = eg.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start sandboxes: %v\n", err)
	}
}

func mockSnapshot(
	ctx context.Context,
	templateId,
	buildId,
	sandboxId string,
	dns *dns.DNS,
	keepAlive time.Duration,
	networkPool *network.Pool,
	templateCache *template.Cache,
) error {
	tracer := otel.Tracer(fmt.Sprintf("sandbox-%s", sandboxId))
	childCtx, _ := tracer.Start(ctx, "mock-sandbox")

	logger := logs.NewSandboxLogger(sandboxId, templateId, "test-team", 2, 512, false)

	start := time.Now()

	sbx, cleanup, err := sandbox.NewSandbox(
		childCtx,
		tracer,
		dns,
		networkPool,
		templateCache,
		&orchestrator.SandboxConfig{
			TemplateId:         templateId,
			FirecrackerVersion: "v1.7.0-dev_8bb88311",
			KernelVersion:      "vmlinux-5.10.186",
			TeamId:             "test-team",
			BuildId:            buildId,
			HugePages:          true,
			MaxSandboxLength:   1,
			SandboxId:          sandboxId,
			EnvdVersion:        "0.1.1",
			RamMb:              512,
			Vcpu:               2,
		},
		"trace-test-1",
		time.Now(),
		time.Now(),
		logger,
	)
	defer func() {
		cleanupErr := cleanup.Run()
		if cleanupErr != nil {
			fmt.Fprintf(os.Stderr, "failed to cleanup sandbox: %v\n", cleanupErr)
		}
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create sandbox: %v\n", err)

		return err
	}

	duration := time.Since(start)

	fmt.Printf("[Sandbox is running] - started in %dms \n", duration.Milliseconds())

	time.Sleep(keepAlive)

	fmt.Println("Snapshotting sandbox")

	snapshotTime := time.Now()

	snapshotTemplateFiles, err := storage.NewTemplateFiles(
		"snapshot-template",
		"snapshot-build",
		sbx.Config.KernelVersion,
		sbx.Config.FirecrackerVersion,
		sbx.Config.HugePages,
	).NewTemplateCacheFiles()
	if err != nil {
		return fmt.Errorf("failed to create snapshot template files: %w", err)
	}

	err = os.MkdirAll(snapshotTemplateFiles.CacheDir(), 0o755)
	if err != nil {
		return fmt.Errorf("failed to create snapshot template files directory: %w", err)
	}

	err = sbx.Snapshot(ctx, snapshotTemplateFiles)
	if err != nil {
		return fmt.Errorf("failed to snapshot sandbox: %w", err)
	}

	_, err = templateCache.AddTemplateFromLocalFiles(snapshotTemplateFiles)
	if err != nil {
		return fmt.Errorf("failed to add template from local files: %w", err)
	}

	duration = time.Since(snapshotTime)

	fmt.Printf("[Sandbox snapshot] - took %dms \n", duration.Milliseconds())

	return nil
}
