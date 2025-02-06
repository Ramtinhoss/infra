package handlers

import (
	"log"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/e2b-dev/infra/packages/api/internal/api"
	"github.com/e2b-dev/infra/packages/api/internal/auth"
	authcache "github.com/e2b-dev/infra/packages/api/internal/cache/auth"
	"github.com/e2b-dev/infra/packages/shared/pkg/models"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/envbuild"
	"github.com/e2b-dev/infra/packages/shared/pkg/telemetry"
)

func (a *APIStore) GetSandboxes(c *gin.Context, params api.GetSandboxesParams) {
	ctx := c.Request.Context()
	teamInfo := c.Value(auth.TeamContextKey).(authcache.AuthTeamInfo)
	team := teamInfo.Team

	telemetry.ReportEvent(ctx, "list running instances")

	// Initialize empty slice for results
	sandboxes := make([]api.ListedSandbox, 0)

	// Only fetch running instances if we need them (state is nil or "running")
	if params.State == nil || *params.State == "running" {
		instanceInfo := a.orchestrator.GetSandboxes(ctx, &team.ID)

		// Get build IDs for running instances
		buildIDs := make([]uuid.UUID, 0)
		for _, info := range instanceInfo {
			if info.TeamID != nil && *info.TeamID == team.ID && info.BuildID != nil {
				buildIDs = append(buildIDs, *info.BuildID)
			}
		}

		// Only fetch builds if we have running instances
		if len(buildIDs) > 0 {
			builds, err := a.db.Client.EnvBuild.Query().Where(envbuild.IDIn(buildIDs...)).All(ctx)
			if err != nil {
				telemetry.ReportCriticalError(ctx, err)
				return
			}

			buildsMap := make(map[uuid.UUID]*models.EnvBuild, len(builds))
			for _, build := range builds {
				buildsMap[build.ID] = build
			}

			// Add running instances to results
			for _, info := range instanceInfo {
				if info.TeamID == nil || *info.TeamID != team.ID || info.BuildID == nil {
					continue
				}

				memoryMB := int32(-1)
				cpuCount := int32(-1)

				if buildsMap[*info.BuildID] != nil {
					memoryMB = int32(buildsMap[*info.BuildID].RAMMB)
					cpuCount = int32(buildsMap[*info.BuildID].Vcpu)
				}

				instance := api.ListedSandbox{
					ClientID:   info.Instance.ClientID,
					TemplateID: info.Instance.TemplateID,
					Alias:      info.Instance.Alias,
					SandboxID:  info.Instance.SandboxID,
					StartedAt:  info.StartTime,
					CpuCount:   cpuCount,
					MemoryMB:   memoryMB,
					EndAt:      info.EndTime,
					State:      "running",
				}

				if info.Metadata != nil {
					meta := api.SandboxMetadata(info.Metadata)
					instance.Metadata = &meta
				}

				sandboxes = append(sandboxes, instance)
			}
		}
	}

	// Only fetch snapshots if we need them (state is nil or "paused")
	if params.State == nil || *params.State == "paused" {
		snapshotEnvs, err := a.db.GetTeamSnapshots(ctx, team.ID)
		if err != nil {
			telemetry.ReportCriticalError(ctx, err)
			return
		}

		// Add snapshots to results
		for _, e := range snapshotEnvs {
			snapshotBuilds := e.Edges.Builds
			snapshot := e.Edges.Snapshots[0]

			log.Printf("snapshot %s, build count %d", snapshot.ID, len(snapshotBuilds))

			memoryMB := int32(-1)
			cpuCount := int32(-1)

			if len(snapshotBuilds) > 0 {
				memoryMB = int32(snapshotBuilds[0].RAMMB)
				cpuCount = int32(snapshotBuilds[0].Vcpu)
			}

			instance := api.ListedSandbox{
				ClientID:   "00000000",
				TemplateID: e.ID,
				SandboxID:  snapshot.SandboxID,
				StartedAt:  snapshot.SandboxStartedAt,
				CpuCount:   cpuCount,
				MemoryMB:   memoryMB,
				EndAt:      snapshot.CreatedAt,
				State:      "paused",
			}

			if snapshot.Metadata != nil {
				meta := api.SandboxMetadata(snapshot.Metadata)
				instance.Metadata = &meta
			}

			sandboxes = append(sandboxes, instance)
		}
	}

	// filter sandboxes by metadata
	if params.Query != nil {
		// Unescape query
		query, err := url.QueryUnescape(*params.Query)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Error when unescaping query")

			return
		}

		// Parse filters, both key and value are also unescaped
		filters := make(map[string]string)

		for _, filter := range strings.Split(query, "&") {
			parts := strings.Split(filter, "=")
			if len(parts) != 2 {
				c.JSON(http.StatusBadRequest, "Invalid key value pair in query")

				return
			}

			key, err := url.QueryUnescape(parts[0])
			if err != nil {
				c.JSON(http.StatusBadRequest, "Error when unescaping key")

				return
			}

			value, err := url.QueryUnescape(parts[1])
			if err != nil {
				c.JSON(http.StatusBadRequest, "Error when unescaping value")

				return
			}

			filters[key] = value
		}

		// Filter instances to match all filters
		n := 0
		for _, instance := range sandboxes {
			if instance.Metadata == nil {
				continue
			}

			matchesAll := true
			for key, value := range filters {
				if metadataValue, ok := (*instance.Metadata)[key]; !ok || metadataValue != value {
					matchesAll = false
					break
				}
			}

			if matchesAll {
				sandboxes[n] = instance
				n++
			}
		}

		// Trim slice
		sandboxes = sandboxes[:n]
	}

	// Sort sandboxes by start time descending
	slices.SortFunc(sandboxes, func(a, b api.ListedSandbox) int {
		return a.StartedAt.Compare(b.StartedAt)
	})

	// Report analytics
	a.posthog.IdentifyAnalyticsTeam(team.ID.String(), team.Name)
	properties := a.posthog.GetPackageToPosthogProperties(&c.Request.Header)
	a.posthog.CreateAnalyticsTeamEvent(team.ID.String(), "listed running instances", properties)

	c.JSON(http.StatusOK, sandboxes)
}
