// Code generated by ent, DO NOT EDIT.

package usersteams

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/internal"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldEQ(FieldUserID, v))
}

// TeamID applies equality check predicate on the "team_id" field. It's identical to TeamIDEQ.
func TeamID(v uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldEQ(FieldTeamID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldNotIn(FieldUserID, vs...))
}

// TeamIDEQ applies the EQ predicate on the "team_id" field.
func TeamIDEQ(v uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldEQ(FieldTeamID, v))
}

// TeamIDNEQ applies the NEQ predicate on the "team_id" field.
func TeamIDNEQ(v uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldNEQ(FieldTeamID, v))
}

// TeamIDIn applies the In predicate on the "team_id" field.
func TeamIDIn(vs ...uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldIn(FieldTeamID, vs...))
}

// TeamIDNotIn applies the NotIn predicate on the "team_id" field.
func TeamIDNotIn(vs ...uuid.UUID) predicate.UsersTeams {
	return predicate.UsersTeams(sql.FieldNotIn(FieldTeamID, vs...))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.UsersTeams {
	return predicate.UsersTeams(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UsersTable, UsersColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UsersTeams
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.UsersTeams {
	return predicate.UsersTeams(func(s *sql.Selector) {
		step := newUsersStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UsersTeams
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.UsersTeams {
	return predicate.UsersTeams(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TeamsTable, TeamsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.UsersTeams
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.Team) predicate.UsersTeams {
	return predicate.UsersTeams(func(s *sql.Selector) {
		step := newTeamsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.UsersTeams
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UsersTeams) predicate.UsersTeams {
	return predicate.UsersTeams(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UsersTeams) predicate.UsersTeams {
	return predicate.UsersTeams(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UsersTeams) predicate.UsersTeams {
	return predicate.UsersTeams(sql.NotPredicates(p))
}
