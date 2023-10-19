// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/env"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/internal"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/predicate"
	"github.com/e2b-dev/infra/packages/api/internal/db/ent/team"
)

// EnvQuery is the builder for querying Env entities.
type EnvQuery struct {
	config
	ctx        *QueryContext
	order      []env.OrderOption
	inters     []Interceptor
	predicates []predicate.Env
	withTeam   *TeamQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EnvQuery builder.
func (eq *EnvQuery) Where(ps ...predicate.Env) *EnvQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EnvQuery) Limit(limit int) *EnvQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EnvQuery) Offset(offset int) *EnvQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EnvQuery) Unique(unique bool) *EnvQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EnvQuery) Order(o ...env.OrderOption) *EnvQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryTeam chains the current query on the "team" edge.
func (eq *EnvQuery) QueryTeam() *TeamQuery {
	query := (&TeamClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(env.Table, env.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, env.TeamTable, env.TeamColumn),
		)
		schemaConfig := eq.schemaConfig
		step.To.Schema = schemaConfig.Team
		step.Edge.Schema = schemaConfig.Team
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Env entity from the query.
// Returns a *NotFoundError when no Env was found.
func (eq *EnvQuery) First(ctx context.Context) (*Env, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{env.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EnvQuery) FirstX(ctx context.Context) *Env {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Env ID from the query.
// Returns a *NotFoundError when no Env ID was found.
func (eq *EnvQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{env.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EnvQuery) FirstIDX(ctx context.Context) string {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Env entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Env entity is found.
// Returns a *NotFoundError when no Env entities are found.
func (eq *EnvQuery) Only(ctx context.Context) (*Env, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{env.Label}
	default:
		return nil, &NotSingularError{env.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EnvQuery) OnlyX(ctx context.Context) *Env {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Env ID in the query.
// Returns a *NotSingularError when more than one Env ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EnvQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{env.Label}
	default:
		err = &NotSingularError{env.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EnvQuery) OnlyIDX(ctx context.Context) string {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Envs.
func (eq *EnvQuery) All(ctx context.Context) ([]*Env, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Env, *EnvQuery]()
	return withInterceptors[[]*Env](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EnvQuery) AllX(ctx context.Context) []*Env {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Env IDs.
func (eq *EnvQuery) IDs(ctx context.Context) (ids []string, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(env.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EnvQuery) IDsX(ctx context.Context) []string {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EnvQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EnvQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EnvQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EnvQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EnvQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EnvQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EnvQuery) Clone() *EnvQuery {
	if eq == nil {
		return nil
	}
	return &EnvQuery{
		config:     eq.config,
		ctx:        eq.ctx.Clone(),
		order:      append([]env.OrderOption{}, eq.order...),
		inters:     append([]Interceptor{}, eq.inters...),
		predicates: append([]predicate.Env{}, eq.predicates...),
		withTeam:   eq.withTeam.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithTeam tells the query-builder to eager-load the nodes that are connected to
// the "team" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EnvQuery) WithTeam(opts ...func(*TeamQuery)) *EnvQuery {
	query := (&TeamClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withTeam = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Env.Query().
//		GroupBy(env.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EnvQuery) GroupBy(field string, fields ...string) *EnvGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EnvGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = env.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Env.Query().
//		Select(env.FieldCreatedAt).
//		Scan(ctx, &v)
func (eq *EnvQuery) Select(fields ...string) *EnvSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EnvSelect{EnvQuery: eq}
	sbuild.label = env.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EnvSelect configured with the given aggregations.
func (eq *EnvQuery) Aggregate(fns ...AggregateFunc) *EnvSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EnvQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !env.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EnvQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Env, error) {
	var (
		nodes       = []*Env{}
		_spec       = eq.querySpec()
		loadedTypes = [1]bool{
			eq.withTeam != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Env).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Env{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = eq.schemaConfig.Env
	ctx = internal.NewSchemaConfigContext(ctx, eq.schemaConfig)
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withTeam; query != nil {
		if err := eq.loadTeam(ctx, query, nodes,
			func(n *Env) { n.Edges.Team = []*Team{} },
			func(n *Env, e *Team) { n.Edges.Team = append(n.Edges.Team, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EnvQuery) loadTeam(ctx context.Context, query *TeamQuery, nodes []*Env, init func(*Env), assign func(*Env, *Team)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Env)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Team(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(env.TeamColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.env_team
		if fk == nil {
			return fmt.Errorf(`foreign-key "env_team" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "env_team" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eq *EnvQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Schema = eq.schemaConfig.Env
	ctx = internal.NewSchemaConfigContext(ctx, eq.schemaConfig)
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EnvQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(env.Table, env.Columns, sqlgraph.NewFieldSpec(env.FieldID, field.TypeString))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, env.FieldID)
		for i := range fields {
			if fields[i] != env.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EnvQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(env.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = env.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(eq.schemaConfig.Env)
	ctx = internal.NewSchemaConfigContext(ctx, eq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EnvGroupBy is the group-by builder for Env entities.
type EnvGroupBy struct {
	selector
	build *EnvQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EnvGroupBy) Aggregate(fns ...AggregateFunc) *EnvGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EnvGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnvQuery, *EnvGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EnvGroupBy) sqlScan(ctx context.Context, root *EnvQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EnvSelect is the builder for selecting fields of Env entities.
type EnvSelect struct {
	*EnvQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EnvSelect) Aggregate(fns ...AggregateFunc) *EnvSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EnvSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EnvQuery, *EnvSelect](ctx, es.EnvQuery, es, es.inters, v)
}

func (es *EnvSelect) sqlScan(ctx context.Context, root *EnvQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
