// Code generated by ent, DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/env"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/team"
	"github.com/google/uuid"
)

// Env is the model entity for the Env schema.
type Env struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TeamID holds the value of the "team_id" field.
	TeamID uuid.UUID `json:"team_id,omitempty"`
	// Dockerfile holds the value of the "dockerfile" field.
	Dockerfile string `json:"dockerfile,omitempty"`
	// Public holds the value of the "public" field.
	Public bool `json:"public,omitempty"`
	// BuildID holds the value of the "build_id" field.
	BuildID uuid.UUID `json:"build_id,omitempty"`
	// BuildCount holds the value of the "build_count" field.
	BuildCount int32 `json:"build_count,omitempty"`
	// SpawnCount holds the value of the "spawn_count" field.
	SpawnCount int32 `json:"spawn_count,omitempty"`
	// LastSpawnedAt holds the value of the "last_spawned_at" field.
	LastSpawnedAt time.Time `json:"last_spawned_at,omitempty"`
	// Vcpu holds the value of the "vcpu" field.
	Vcpu int64 `json:"vcpu,omitempty"`
	// RAMMB holds the value of the "ram_mb" field.
	RAMMB int64 `json:"ram_mb,omitempty"`
	// FreeDiskSizeMB holds the value of the "free_disk_size_mb" field.
	FreeDiskSizeMB int64 `json:"free_disk_size_mb,omitempty"`
	// TotalDiskSizeMB holds the value of the "total_disk_size_mb" field.
	TotalDiskSizeMB int64 `json:"total_disk_size_mb,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EnvQuery when eager-loading is set.
	Edges        EnvEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EnvEdges holds the relations/edges for other nodes in the graph.
type EnvEdges struct {
	// Team holds the value of the team edge.
	Team *Team `json:"team,omitempty"`
	// EnvAliases holds the value of the env_aliases edge.
	EnvAliases []*EnvAlias `json:"env_aliases,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EnvEdges) TeamOrErr() (*Team, error) {
	if e.loadedTypes[0] {
		if e.Team == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: team.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// EnvAliasesOrErr returns the EnvAliases value or an error if the edge
// was not loaded in eager-loading.
func (e EnvEdges) EnvAliasesOrErr() ([]*EnvAlias, error) {
	if e.loadedTypes[1] {
		return e.EnvAliases, nil
	}
	return nil, &NotLoadedError{edge: "env_aliases"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Env) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case env.FieldPublic:
			values[i] = new(sql.NullBool)
		case env.FieldBuildCount, env.FieldSpawnCount, env.FieldVcpu, env.FieldRAMMB, env.FieldFreeDiskSizeMB, env.FieldTotalDiskSizeMB:
			values[i] = new(sql.NullInt64)
		case env.FieldID, env.FieldDockerfile:
			values[i] = new(sql.NullString)
		case env.FieldCreatedAt, env.FieldUpdatedAt, env.FieldLastSpawnedAt:
			values[i] = new(sql.NullTime)
		case env.FieldTeamID, env.FieldBuildID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Env fields.
func (e *Env) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case env.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = value.String
			}
		case env.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case env.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case env.FieldTeamID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value != nil {
				e.TeamID = *value
			}
		case env.FieldDockerfile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dockerfile", values[i])
			} else if value.Valid {
				e.Dockerfile = value.String
			}
		case env.FieldPublic:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field public", values[i])
			} else if value.Valid {
				e.Public = value.Bool
			}
		case env.FieldBuildID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field build_id", values[i])
			} else if value != nil {
				e.BuildID = *value
			}
		case env.FieldBuildCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field build_count", values[i])
			} else if value.Valid {
				e.BuildCount = int32(value.Int64)
			}
		case env.FieldSpawnCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field spawn_count", values[i])
			} else if value.Valid {
				e.SpawnCount = int32(value.Int64)
			}
		case env.FieldLastSpawnedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_spawned_at", values[i])
			} else if value.Valid {
				e.LastSpawnedAt = value.Time
			}
		case env.FieldVcpu:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vcpu", values[i])
			} else if value.Valid {
				e.Vcpu = value.Int64
			}
		case env.FieldRAMMB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ram_mb", values[i])
			} else if value.Valid {
				e.RAMMB = value.Int64
			}
		case env.FieldFreeDiskSizeMB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field free_disk_size_mb", values[i])
			} else if value.Valid {
				e.FreeDiskSizeMB = value.Int64
			}
		case env.FieldTotalDiskSizeMB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_disk_size_mb", values[i])
			} else if value.Valid {
				e.TotalDiskSizeMB = value.Int64
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Env.
// This includes values selected through modifiers, order, etc.
func (e *Env) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryTeam queries the "team" edge of the Env entity.
func (e *Env) QueryTeam() *TeamQuery {
	return NewEnvClient(e.config).QueryTeam(e)
}

// QueryEnvAliases queries the "env_aliases" edge of the Env entity.
func (e *Env) QueryEnvAliases() *EnvAliasQuery {
	return NewEnvClient(e.config).QueryEnvAliases(e)
}

// Update returns a builder for updating this Env.
// Note that you need to call Env.Unwrap() before calling this method if this Env
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Env) Update() *EnvUpdateOne {
	return NewEnvClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Env entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Env) Unwrap() *Env {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("models: Env is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Env) String() string {
	var builder strings.Builder
	builder.WriteString("Env(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(fmt.Sprintf("%v", e.TeamID))
	builder.WriteString(", ")
	builder.WriteString("dockerfile=")
	builder.WriteString(e.Dockerfile)
	builder.WriteString(", ")
	builder.WriteString("public=")
	builder.WriteString(fmt.Sprintf("%v", e.Public))
	builder.WriteString(", ")
	builder.WriteString("build_id=")
	builder.WriteString(fmt.Sprintf("%v", e.BuildID))
	builder.WriteString(", ")
	builder.WriteString("build_count=")
	builder.WriteString(fmt.Sprintf("%v", e.BuildCount))
	builder.WriteString(", ")
	builder.WriteString("spawn_count=")
	builder.WriteString(fmt.Sprintf("%v", e.SpawnCount))
	builder.WriteString(", ")
	builder.WriteString("last_spawned_at=")
	builder.WriteString(e.LastSpawnedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("vcpu=")
	builder.WriteString(fmt.Sprintf("%v", e.Vcpu))
	builder.WriteString(", ")
	builder.WriteString("ram_mb=")
	builder.WriteString(fmt.Sprintf("%v", e.RAMMB))
	builder.WriteString(", ")
	builder.WriteString("free_disk_size_mb=")
	builder.WriteString(fmt.Sprintf("%v", e.FreeDiskSizeMB))
	builder.WriteString(", ")
	builder.WriteString("total_disk_size_mb=")
	builder.WriteString(fmt.Sprintf("%v", e.TotalDiskSizeMB))
	builder.WriteByte(')')
	return builder.String()
}

// Envs is a parsable slice of Env.
type Envs []*Env
