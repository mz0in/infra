// Code generated by ent, DO NOT EDIT.

package teamapikey

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/internal"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldEQ(FieldCreatedAt, v))
}

// TeamID applies equality check predicate on the "team_id" field. It's identical to TeamIDEQ.
func TeamID(v uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldEQ(FieldTeamID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldLTE(FieldCreatedAt, v))
}

// TeamIDEQ applies the EQ predicate on the "team_id" field.
func TeamIDEQ(v uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldEQ(FieldTeamID, v))
}

// TeamIDNEQ applies the NEQ predicate on the "team_id" field.
func TeamIDNEQ(v uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldNEQ(FieldTeamID, v))
}

// TeamIDIn applies the In predicate on the "team_id" field.
func TeamIDIn(vs ...uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldIn(FieldTeamID, vs...))
}

// TeamIDNotIn applies the NotIn predicate on the "team_id" field.
func TeamIDNotIn(vs ...uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldNotIn(FieldTeamID, vs...))
}

// TeamIDGT applies the GT predicate on the "team_id" field.
func TeamIDGT(v uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldGT(FieldTeamID, v))
}

// TeamIDGTE applies the GTE predicate on the "team_id" field.
func TeamIDGTE(v uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldGTE(FieldTeamID, v))
}

// TeamIDLT applies the LT predicate on the "team_id" field.
func TeamIDLT(v uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldLT(FieldTeamID, v))
}

// TeamIDLTE applies the LTE predicate on the "team_id" field.
func TeamIDLTE(v uuid.UUID) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.FieldLTE(FieldTeamID, v))
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.TeamApiKey {
	return predicate.TeamApiKey(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeamTable, TeamColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.Team
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Team) predicate.TeamApiKey {
	return predicate.TeamApiKey(func(s *sql.Selector) {
		step := newTeamStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.Team
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TeamApiKey) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TeamApiKey) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TeamApiKey) predicate.TeamApiKey {
	return predicate.TeamApiKey(sql.NotPredicates(p))
}
