// Code generated by ent, DO NOT EDIT.

package team

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldIsDefault holds the string denoting the is_default field in the database.
	FieldIsDefault = "is_default"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTier holds the string denoting the tier field in the database.
	FieldTier = "tier"
	// FieldIsBlocked holds the string denoting the is_blocked field in the database.
	FieldIsBlocked = "is_blocked"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeTeamAPIKeys holds the string denoting the team_api_keys edge name in mutations.
	EdgeTeamAPIKeys = "team_api_keys"
	// EdgeTeamTier holds the string denoting the team_tier edge name in mutations.
	EdgeTeamTier = "team_tier"
	// EdgeEnvs holds the string denoting the envs edge name in mutations.
	EdgeEnvs = "envs"
	// EdgeUsersTeams holds the string denoting the users_teams edge name in mutations.
	EdgeUsersTeams = "users_teams"
	// TeamApiKeyFieldID holds the string denoting the ID field of the TeamApiKey.
	TeamApiKeyFieldID = "api_key"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "users_teams"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// TeamAPIKeysTable is the table that holds the team_api_keys relation/edge.
	TeamAPIKeysTable = "team_api_keys"
	// TeamAPIKeysInverseTable is the table name for the TeamApiKey entity.
	// It exists in this package in order to avoid circular dependency with the "teamapikey" package.
	TeamAPIKeysInverseTable = "team_api_keys"
	// TeamAPIKeysColumn is the table column denoting the team_api_keys relation/edge.
	TeamAPIKeysColumn = "team_id"
	// TeamTierTable is the table that holds the team_tier relation/edge.
	TeamTierTable = "teams"
	// TeamTierInverseTable is the table name for the Tier entity.
	// It exists in this package in order to avoid circular dependency with the "tier" package.
	TeamTierInverseTable = "tiers"
	// TeamTierColumn is the table column denoting the team_tier relation/edge.
	TeamTierColumn = "tier"
	// EnvsTable is the table that holds the envs relation/edge.
	EnvsTable = "envs"
	// EnvsInverseTable is the table name for the Env entity.
	// It exists in this package in order to avoid circular dependency with the "env" package.
	EnvsInverseTable = "envs"
	// EnvsColumn is the table column denoting the envs relation/edge.
	EnvsColumn = "team_id"
	// UsersTeamsTable is the table that holds the users_teams relation/edge.
	UsersTeamsTable = "users_teams"
	// UsersTeamsInverseTable is the table name for the UsersTeams entity.
	// It exists in this package in order to avoid circular dependency with the "usersteams" package.
	UsersTeamsInverseTable = "users_teams"
	// UsersTeamsColumn is the table column denoting the users_teams relation/edge.
	UsersTeamsColumn = "team_id"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldIsDefault,
	FieldName,
	FieldTier,
	FieldIsBlocked,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"team_id", "user_id"}
)

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

// OrderOption defines the ordering options for the Team queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByIsDefault orders the results by the is_default field.
func ByIsDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDefault, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTier orders the results by the tier field.
func ByTier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTier, opts...).ToFunc()
}

// ByIsBlocked orders the results by the is_blocked field.
func ByIsBlocked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsBlocked, opts...).ToFunc()
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTeamAPIKeysCount orders the results by team_api_keys count.
func ByTeamAPIKeysCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamAPIKeysStep(), opts...)
	}
}

// ByTeamAPIKeys orders the results by team_api_keys terms.
func ByTeamAPIKeys(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamAPIKeysStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTeamTierField orders the results by team_tier field.
func ByTeamTierField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamTierStep(), sql.OrderByField(field, opts...))
	}
}

// ByEnvsCount orders the results by envs count.
func ByEnvsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEnvsStep(), opts...)
	}
}

// ByEnvs orders the results by envs terms.
func ByEnvs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEnvsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUsersTeamsCount orders the results by users_teams count.
func ByUsersTeamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersTeamsStep(), opts...)
	}
}

// ByUsersTeams orders the results by users_teams terms.
func ByUsersTeams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersTeamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
	)
}
func newTeamAPIKeysStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamAPIKeysInverseTable, TeamApiKeyFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TeamAPIKeysTable, TeamAPIKeysColumn),
	)
}
func newTeamTierStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamTierInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TeamTierTable, TeamTierColumn),
	)
}
func newEnvsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EnvsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EnvsTable, EnvsColumn),
	)
}
func newUsersTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersTeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UsersTeamsTable, UsersTeamsColumn),
	)
}
