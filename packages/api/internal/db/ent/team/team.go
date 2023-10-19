// Code generated by ent, DO NOT EDIT.

package team

import (
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
	// FieldIsBlocked holds the string denoting the is_blocked field in the database.
	FieldIsBlocked = "is_blocked"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeUsersTeams holds the string denoting the users_teams edge name in mutations.
	EdgeUsersTeams = "users_teams"
	// Table holds the table name of the team in the database.
	Table = "teams"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "users_teams"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
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
	FieldIsBlocked,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "teams"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"env_team",
	"team_api_key_team",
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
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

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
func newUsersTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersTeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UsersTeamsTable, UsersTeamsColumn),
	)
}
