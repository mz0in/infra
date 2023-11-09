// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/env"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/envalias"
)

// EnvAliasCreate is the builder for creating a EnvAlias entity.
type EnvAliasCreate struct {
	config
	mutation *EnvAliasMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEnvID sets the "env_id" field.
func (eac *EnvAliasCreate) SetEnvID(s string) *EnvAliasCreate {
	eac.mutation.SetEnvID(s)
	return eac
}

// SetNillableEnvID sets the "env_id" field if the given value is not nil.
func (eac *EnvAliasCreate) SetNillableEnvID(s *string) *EnvAliasCreate {
	if s != nil {
		eac.SetEnvID(*s)
	}
	return eac
}

// SetIsName sets the "is_name" field.
func (eac *EnvAliasCreate) SetIsName(b bool) *EnvAliasCreate {
	eac.mutation.SetIsName(b)
	return eac
}

// SetID sets the "id" field.
func (eac *EnvAliasCreate) SetID(s string) *EnvAliasCreate {
	eac.mutation.SetID(s)
	return eac
}

// SetAliasEnvID sets the "alias_env" edge to the Env entity by ID.
func (eac *EnvAliasCreate) SetAliasEnvID(id string) *EnvAliasCreate {
	eac.mutation.SetAliasEnvID(id)
	return eac
}

// SetNillableAliasEnvID sets the "alias_env" edge to the Env entity by ID if the given value is not nil.
func (eac *EnvAliasCreate) SetNillableAliasEnvID(id *string) *EnvAliasCreate {
	if id != nil {
		eac = eac.SetAliasEnvID(*id)
	}
	return eac
}

// SetAliasEnv sets the "alias_env" edge to the Env entity.
func (eac *EnvAliasCreate) SetAliasEnv(e *Env) *EnvAliasCreate {
	return eac.SetAliasEnvID(e.ID)
}

// Mutation returns the EnvAliasMutation object of the builder.
func (eac *EnvAliasCreate) Mutation() *EnvAliasMutation {
	return eac.mutation
}

// Save creates the EnvAlias in the database.
func (eac *EnvAliasCreate) Save(ctx context.Context) (*EnvAlias, error) {
	return withHooks(ctx, eac.sqlSave, eac.mutation, eac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eac *EnvAliasCreate) SaveX(ctx context.Context) *EnvAlias {
	v, err := eac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eac *EnvAliasCreate) Exec(ctx context.Context) error {
	_, err := eac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eac *EnvAliasCreate) ExecX(ctx context.Context) {
	if err := eac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eac *EnvAliasCreate) check() error {
	if _, ok := eac.mutation.IsName(); !ok {
		return &ValidationError{Name: "is_name", err: errors.New(`ent: missing required field "EnvAlias.is_name"`)}
	}
	return nil
}

func (eac *EnvAliasCreate) sqlSave(ctx context.Context) (*EnvAlias, error) {
	if err := eac.check(); err != nil {
		return nil, err
	}
	_node, _spec := eac.createSpec()
	if err := sqlgraph.CreateNode(ctx, eac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected EnvAlias.ID type: %T", _spec.ID.Value)
		}
	}
	eac.mutation.id = &_node.ID
	eac.mutation.done = true
	return _node, nil
}

func (eac *EnvAliasCreate) createSpec() (*EnvAlias, *sqlgraph.CreateSpec) {
	var (
		_node = &EnvAlias{config: eac.config}
		_spec = sqlgraph.NewCreateSpec(envalias.Table, sqlgraph.NewFieldSpec(envalias.FieldID, field.TypeString))
	)
	_spec.Schema = eac.schemaConfig.EnvAlias
	_spec.OnConflict = eac.conflict
	if id, ok := eac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := eac.mutation.IsName(); ok {
		_spec.SetField(envalias.FieldIsName, field.TypeBool, value)
		_node.IsName = value
	}
	if nodes := eac.mutation.AliasEnvIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   envalias.AliasEnvTable,
			Columns: []string{envalias.AliasEnvColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = eac.schemaConfig.EnvAlias
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EnvID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EnvAlias.Create().
//		SetEnvID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EnvAliasUpsert) {
//			SetEnvID(v+v).
//		}).
//		Exec(ctx)
func (eac *EnvAliasCreate) OnConflict(opts ...sql.ConflictOption) *EnvAliasUpsertOne {
	eac.conflict = opts
	return &EnvAliasUpsertOne{
		create: eac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EnvAlias.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (eac *EnvAliasCreate) OnConflictColumns(columns ...string) *EnvAliasUpsertOne {
	eac.conflict = append(eac.conflict, sql.ConflictColumns(columns...))
	return &EnvAliasUpsertOne{
		create: eac,
	}
}

type (
	// EnvAliasUpsertOne is the builder for "upsert"-ing
	//  one EnvAlias node.
	EnvAliasUpsertOne struct {
		create *EnvAliasCreate
	}

	// EnvAliasUpsert is the "OnConflict" setter.
	EnvAliasUpsert struct {
		*sql.UpdateSet
	}
)

// SetEnvID sets the "env_id" field.
func (u *EnvAliasUpsert) SetEnvID(v string) *EnvAliasUpsert {
	u.Set(envalias.FieldEnvID, v)
	return u
}

// UpdateEnvID sets the "env_id" field to the value that was provided on create.
func (u *EnvAliasUpsert) UpdateEnvID() *EnvAliasUpsert {
	u.SetExcluded(envalias.FieldEnvID)
	return u
}

// ClearEnvID clears the value of the "env_id" field.
func (u *EnvAliasUpsert) ClearEnvID() *EnvAliasUpsert {
	u.SetNull(envalias.FieldEnvID)
	return u
}

// SetIsName sets the "is_name" field.
func (u *EnvAliasUpsert) SetIsName(v bool) *EnvAliasUpsert {
	u.Set(envalias.FieldIsName, v)
	return u
}

// UpdateIsName sets the "is_name" field to the value that was provided on create.
func (u *EnvAliasUpsert) UpdateIsName() *EnvAliasUpsert {
	u.SetExcluded(envalias.FieldIsName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.EnvAlias.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(envalias.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EnvAliasUpsertOne) UpdateNewValues() *EnvAliasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(envalias.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EnvAlias.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EnvAliasUpsertOne) Ignore() *EnvAliasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EnvAliasUpsertOne) DoNothing() *EnvAliasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EnvAliasCreate.OnConflict
// documentation for more info.
func (u *EnvAliasUpsertOne) Update(set func(*EnvAliasUpsert)) *EnvAliasUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EnvAliasUpsert{UpdateSet: update})
	}))
	return u
}

// SetEnvID sets the "env_id" field.
func (u *EnvAliasUpsertOne) SetEnvID(v string) *EnvAliasUpsertOne {
	return u.Update(func(s *EnvAliasUpsert) {
		s.SetEnvID(v)
	})
}

// UpdateEnvID sets the "env_id" field to the value that was provided on create.
func (u *EnvAliasUpsertOne) UpdateEnvID() *EnvAliasUpsertOne {
	return u.Update(func(s *EnvAliasUpsert) {
		s.UpdateEnvID()
	})
}

// ClearEnvID clears the value of the "env_id" field.
func (u *EnvAliasUpsertOne) ClearEnvID() *EnvAliasUpsertOne {
	return u.Update(func(s *EnvAliasUpsert) {
		s.ClearEnvID()
	})
}

// SetIsName sets the "is_name" field.
func (u *EnvAliasUpsertOne) SetIsName(v bool) *EnvAliasUpsertOne {
	return u.Update(func(s *EnvAliasUpsert) {
		s.SetIsName(v)
	})
}

// UpdateIsName sets the "is_name" field to the value that was provided on create.
func (u *EnvAliasUpsertOne) UpdateIsName() *EnvAliasUpsertOne {
	return u.Update(func(s *EnvAliasUpsert) {
		s.UpdateIsName()
	})
}

// Exec executes the query.
func (u *EnvAliasUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EnvAliasCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EnvAliasUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EnvAliasUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: EnvAliasUpsertOne.ID is not supported by MySQL driver. Use EnvAliasUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EnvAliasUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EnvAliasCreateBulk is the builder for creating many EnvAlias entities in bulk.
type EnvAliasCreateBulk struct {
	config
	err      error
	builders []*EnvAliasCreate
	conflict []sql.ConflictOption
}

// Save creates the EnvAlias entities in the database.
func (eacb *EnvAliasCreateBulk) Save(ctx context.Context) ([]*EnvAlias, error) {
	if eacb.err != nil {
		return nil, eacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(eacb.builders))
	nodes := make([]*EnvAlias, len(eacb.builders))
	mutators := make([]Mutator, len(eacb.builders))
	for i := range eacb.builders {
		func(i int, root context.Context) {
			builder := eacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EnvAliasMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, eacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = eacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eacb *EnvAliasCreateBulk) SaveX(ctx context.Context) []*EnvAlias {
	v, err := eacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eacb *EnvAliasCreateBulk) Exec(ctx context.Context) error {
	_, err := eacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eacb *EnvAliasCreateBulk) ExecX(ctx context.Context) {
	if err := eacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EnvAlias.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EnvAliasUpsert) {
//			SetEnvID(v+v).
//		}).
//		Exec(ctx)
func (eacb *EnvAliasCreateBulk) OnConflict(opts ...sql.ConflictOption) *EnvAliasUpsertBulk {
	eacb.conflict = opts
	return &EnvAliasUpsertBulk{
		create: eacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EnvAlias.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (eacb *EnvAliasCreateBulk) OnConflictColumns(columns ...string) *EnvAliasUpsertBulk {
	eacb.conflict = append(eacb.conflict, sql.ConflictColumns(columns...))
	return &EnvAliasUpsertBulk{
		create: eacb,
	}
}

// EnvAliasUpsertBulk is the builder for "upsert"-ing
// a bulk of EnvAlias nodes.
type EnvAliasUpsertBulk struct {
	create *EnvAliasCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.EnvAlias.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(envalias.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EnvAliasUpsertBulk) UpdateNewValues() *EnvAliasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(envalias.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EnvAlias.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EnvAliasUpsertBulk) Ignore() *EnvAliasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EnvAliasUpsertBulk) DoNothing() *EnvAliasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EnvAliasCreateBulk.OnConflict
// documentation for more info.
func (u *EnvAliasUpsertBulk) Update(set func(*EnvAliasUpsert)) *EnvAliasUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EnvAliasUpsert{UpdateSet: update})
	}))
	return u
}

// SetEnvID sets the "env_id" field.
func (u *EnvAliasUpsertBulk) SetEnvID(v string) *EnvAliasUpsertBulk {
	return u.Update(func(s *EnvAliasUpsert) {
		s.SetEnvID(v)
	})
}

// UpdateEnvID sets the "env_id" field to the value that was provided on create.
func (u *EnvAliasUpsertBulk) UpdateEnvID() *EnvAliasUpsertBulk {
	return u.Update(func(s *EnvAliasUpsert) {
		s.UpdateEnvID()
	})
}

// ClearEnvID clears the value of the "env_id" field.
func (u *EnvAliasUpsertBulk) ClearEnvID() *EnvAliasUpsertBulk {
	return u.Update(func(s *EnvAliasUpsert) {
		s.ClearEnvID()
	})
}

// SetIsName sets the "is_name" field.
func (u *EnvAliasUpsertBulk) SetIsName(v bool) *EnvAliasUpsertBulk {
	return u.Update(func(s *EnvAliasUpsert) {
		s.SetIsName(v)
	})
}

// UpdateIsName sets the "is_name" field to the value that was provided on create.
func (u *EnvAliasUpsertBulk) UpdateIsName() *EnvAliasUpsertBulk {
	return u.Update(func(s *EnvAliasUpsert) {
		s.UpdateIsName()
	})
}

// Exec executes the query.
func (u *EnvAliasUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EnvAliasCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EnvAliasCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EnvAliasUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
