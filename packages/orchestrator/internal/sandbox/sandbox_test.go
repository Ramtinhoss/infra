package sandbox

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/dns"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/network"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/template"
	"github.com/e2b-dev/infra/packages/shared/pkg/db"
	"github.com/e2b-dev/infra/packages/shared/pkg/grpc/orchestrator"
	"github.com/e2b-dev/infra/packages/shared/pkg/logs"
	"github.com/e2b-dev/infra/packages/shared/pkg/storage"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	// envAlias is the alias of the base template to use for the sandbox
	envAlias = "base"
)

type env struct {
	dnsServer     *dns.DNS
	networkPool   *network.Pool
	templateCache *template.Cache
}

type vars struct {
	templateId string
	buildId    string

	fcVersion     string
	kernelVersion string
	envdVersion   string
}

func prepareEnv(ctx context.Context, _ testing.TB) (*env, error) {
	dnsServer := dns.New()

	templateCache, err := template.NewCache(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create template cache: %w", err)
	}

	networkPool, err := network.NewPool(ctx, 1, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to create network pool: %w", err)
	}

	return &env{
		dnsServer:     dnsServer,
		networkPool:   networkPool,
		templateCache: templateCache,
	}, nil
}

type SandboxTestSuite struct {
	suite.Suite
	env  *env
	vars vars

	parent context.Context
	ctx    context.Context
	cancel context.CancelFunc

	sandboxId string
	teamId    string
	tracer    trace.Tracer
	logger    *logs.SandboxLogger
}

func NewSandboxTestSuite(
	ctx context.Context,
	env *env,
) *SandboxTestSuite {
	sandboxId := "test-sandbox-1"
	teamId := "test-team"

	db, err := db.NewClient()
	if err != nil {
		panic(fmt.Sprintf("failed to create db client: %v", err))
	}
	defer db.Close()

	template, build, err := db.GetEnv(ctx, envAlias)
	if err != nil {
		panic(fmt.Sprintf("failed to get %s env: %v", envAlias, err))
	}

	return &SandboxTestSuite{
		env:       env,
		parent:    ctx,
		sandboxId: sandboxId,
		teamId:    teamId,
		tracer:    otel.Tracer(fmt.Sprintf("sandbox-%s", sandboxId)),
		logger:    logs.NewSandboxLogger(sandboxId, template.TemplateID, teamId, 2, 512, false),
		vars: vars{
			templateId:    template.TemplateID,
			buildId:       template.BuildID,
			fcVersion:     build.FirecrackerVersion,
			kernelVersion: build.KernelVersion,
			envdVersion:   *build.EnvdVersion,
		},
	}
}

func TestSandboxTestSuite(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lEnv, err := prepareEnv(ctx, t)
	if err != nil {
		t.Fatalf("failed to prepare environment: %v", err)
	}

	suite.Run(t, NewSandboxTestSuite(ctx, lEnv))
}

func (suite *SandboxTestSuite) SetupTest() {
	suite.ctx, suite.cancel = context.WithCancel(suite.parent)
}

func (suite *SandboxTestSuite) TearDownTest() {
	if suite.cancel != nil {
		suite.cancel()
	}
}

func (suite *SandboxTestSuite) TestNewSandbox() {
	sbx, cleanup, err := suite.createSandbox(512, 2)
	suite.Require().NoError(err)
	defer cleanup.Run()

	suite.Require().NotNil(sbx)
}

func (suite *SandboxTestSuite) TestSandboxSnapshot() {
	sbx, cleanup, err := suite.createSandbox(512, 2)
	suite.Require().NoError(err)
	defer cleanup.Run()

	snapshot, err := suite.snapshotSandbox(sbx)
	suite.Require().NoError(err)

	suite.Assert().NotNil(snapshot)
}

func genericBenchmarkSandbox(b *testing.B, ramMb, vCpu int64) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lEnv, err := prepareEnv(ctx, b)
	if err != nil {
		b.Fatalf("failed to prepare environment: %v", err)
	}
	suite := NewSandboxTestSuite(ctx, lEnv)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		suite.SetupTest()

		sbx, cleanup, err := suite.createSandbox(ramMb, vCpu)
		if err != nil {
			b.Fatalf("failed to create sandbox: %v", err)
		}

		b.StartTimer()
		_, err = suite.snapshotSandbox(sbx)
		b.StopTimer()

		cleanup.Run()
		suite.TearDownTest()

		if err != nil {
			b.Fatalf("failed to snapshot sandbox: %v", err)
		}
	}
}

func BenchmarkCreate(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lEnv, err := prepareEnv(ctx, b)
	if err != nil {
		b.Fatalf("failed to prepare environment: %v", err)
	}
	suite := NewSandboxTestSuite(ctx, lEnv)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		suite.SetupTest()

		b.StartTimer()
		_, cleanup, err := suite.createSandbox(512, 2)
		b.StopTimer()

		if err != nil {
			b.Fatalf("failed to create sandbox: %v", err)
		}
		cleanup.Run()
		suite.TearDownTest()
	}
}

func BenchmarkSandboxSnapshot(b *testing.B) {
	tests := []struct {
		name  string
		ramMb int64
		vCpu  int64
	}{
		{"512MB", 512, 1},
		{"1GB", 1024, 1},
		{"2GB", 2 * 1024, 1},
		{"4GB", 4 * 1024, 1},
		{"8GB", 8 * 1024, 1},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			genericBenchmarkSandbox(b, tt.ramMb, tt.vCpu)
		})
	}
}

func (suite *SandboxTestSuite) snapshotSandbox(sbx *Sandbox) (*Snapshot, error) {
	snapshotTemplateFiles, err := storage.NewTemplateFiles(
		"snapshot-template",
		"f0370054-b669-eee4-b33b-573d5287c6ef",
		sbx.Config.KernelVersion,
		sbx.Config.FirecrackerVersion,
		sbx.Config.HugePages,
	).NewTemplateCacheFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to create snapshot template files: %s", err)
	}

	err = os.MkdirAll(snapshotTemplateFiles.CacheDir(), 0o755)
	if err != nil {
		return nil, fmt.Errorf("failed to create snapshot template files directory: %s", err)
	}
	defer os.RemoveAll(snapshotTemplateFiles.CacheDir())

	snapshot, err := sbx.Snapshot(suite.ctx, suite.tracer, snapshotTemplateFiles, func() {})
	if err != nil {
		return nil, fmt.Errorf("failed to snapshot sandbox: %s", err)
	}

	return snapshot, nil
}

func (suite *SandboxTestSuite) createSandbox(ramMb, vCpu int64) (*Sandbox, *Cleanup, error) {
	childCtx, _ := suite.tracer.Start(suite.ctx, "mock-sandbox")

	sbx, cleanup, err := NewSandbox(
		childCtx,
		suite.tracer,
		suite.env.dnsServer,
		suite.env.networkPool,
		suite.env.templateCache,
		&orchestrator.SandboxConfig{
			TemplateId:         suite.vars.templateId,
			FirecrackerVersion: suite.vars.fcVersion,
			KernelVersion:      suite.vars.kernelVersion,
			TeamId:             suite.teamId,
			BuildId:            suite.vars.buildId,
			HugePages:          true,
			MaxSandboxLength:   1,
			SandboxId:          suite.sandboxId,
			EnvdVersion:        suite.vars.envdVersion,
			RamMb:              ramMb,
			Vcpu:               vCpu,
		},
		"trace-test-1",
		time.Now(),
		time.Now(),
		suite.logger,
		false,
		suite.vars.templateId,
	)
	if err != nil {
		errCleanup := cleanup.Run()
		return nil, nil, fmt.Errorf("failed to create sandbox: %v", errors.Join(err, errCleanup))
	}

	return sbx, cleanup, nil
}
