// Code generated by ent, DO NOT EDIT.

package envalias

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the envalias type in the database.
	Label = "env_alias"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "alias"
	// FieldEnvID holds the string denoting the env_id field in the database.
	FieldEnvID = "env_id"
	// FieldIsName holds the string denoting the is_name field in the database.
	FieldIsName = "is_name"
	// EdgeEnv holds the string denoting the env edge name in mutations.
	EdgeEnv = "env"
	// EnvFieldID holds the string denoting the ID field of the Env.
	EnvFieldID = "id"
	// Table holds the table name of the envalias in the database.
	Table = "env_aliases"
	// EnvTable is the table that holds the env relation/edge.
	EnvTable = "env_aliases"
	// EnvInverseTable is the table name for the Env entity.
	// It exists in this package in order to avoid circular dependency with the "env" package.
	EnvInverseTable = "envs"
	// EnvColumn is the table column denoting the env relation/edge.
	EnvColumn = "env_id"
)

// Columns holds all SQL columns for envalias fields.
var Columns = []string{
	FieldID,
	FieldEnvID,
	FieldIsName,
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
	// DefaultIsName holds the default value on creation for the "is_name" field.
	DefaultIsName bool
)

// OrderOption defines the ordering options for the EnvAlias queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEnvID orders the results by the env_id field.
func ByEnvID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnvID, opts...).ToFunc()
}

// ByIsName orders the results by the is_name field.
func ByIsName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsName, opts...).ToFunc()
}

// ByEnvField orders the results by env field.
func ByEnvField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEnvStep(), sql.OrderByField(field, opts...))
	}
}
func newEnvStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EnvInverseTable, EnvFieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EnvTable, EnvColumn),
	)
}
