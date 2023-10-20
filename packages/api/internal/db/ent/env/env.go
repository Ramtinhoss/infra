// Code generated by ent, DO NOT EDIT.

package env

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the env type in the database.
	Label = "env"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTeamID holds the string denoting the team_id field in the database.
	FieldTeamID = "team_id"
	// FieldDockerfile holds the string denoting the dockerfile field in the database.
	FieldDockerfile = "dockerfile"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// FieldBuildID holds the string denoting the build_id field in the database.
	FieldBuildID = "build_id"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// Table holds the table name of the env in the database.
	Table = "envs"
	// TeamTable is the table that holds the team relation/edge.
	TeamTable = "teams"
	// TeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamInverseTable = "teams"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "env_team"
)

// Columns holds all SQL columns for env fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldTeamID,
	FieldDockerfile,
	FieldStatus,
	FieldPublic,
	FieldBuildID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusBuilding Status = "building"
	StatusReady    Status = "ready"
	StatusError    Status = "error"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusBuilding, StatusReady, StatusError:
		return nil
	default:
		return fmt.Errorf("env: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Env queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByTeamID orders the results by the team_id field.
func ByTeamID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeamID, opts...).ToFunc()
}

// ByDockerfile orders the results by the dockerfile field.
func ByDockerfile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDockerfile, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPublic orders the results by the public field.
func ByPublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublic, opts...).ToFunc()
}

// ByBuildID orders the results by the build_id field.
func ByBuildID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBuildID, opts...).ToFunc()
}

// ByTeamCount orders the results by team count.
func ByTeamCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamStep(), opts...)
	}
}

// ByTeam orders the results by team terms.
func ByTeam(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TeamTable, TeamColumn),
	)
}
