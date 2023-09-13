// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"persona/libs/database/ent/likeschema"
	"persona/libs/database/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LikeSchemaQuery is the builder for querying LikeSchema entities.
type LikeSchemaQuery struct {
	config
	ctx        *QueryContext
	order      []likeschema.OrderOption
	inters     []Interceptor
	predicates []predicate.LikeSchema
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LikeSchemaQuery builder.
func (lsq *LikeSchemaQuery) Where(ps ...predicate.LikeSchema) *LikeSchemaQuery {
	lsq.predicates = append(lsq.predicates, ps...)
	return lsq
}

// Limit the number of records to be returned by this query.
func (lsq *LikeSchemaQuery) Limit(limit int) *LikeSchemaQuery {
	lsq.ctx.Limit = &limit
	return lsq
}

// Offset to start from.
func (lsq *LikeSchemaQuery) Offset(offset int) *LikeSchemaQuery {
	lsq.ctx.Offset = &offset
	return lsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lsq *LikeSchemaQuery) Unique(unique bool) *LikeSchemaQuery {
	lsq.ctx.Unique = &unique
	return lsq
}

// Order specifies how the records should be ordered.
func (lsq *LikeSchemaQuery) Order(o ...likeschema.OrderOption) *LikeSchemaQuery {
	lsq.order = append(lsq.order, o...)
	return lsq
}

// First returns the first LikeSchema entity from the query.
// Returns a *NotFoundError when no LikeSchema was found.
func (lsq *LikeSchemaQuery) First(ctx context.Context) (*LikeSchema, error) {
	nodes, err := lsq.Limit(1).All(setContextOp(ctx, lsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{likeschema.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lsq *LikeSchemaQuery) FirstX(ctx context.Context) *LikeSchema {
	node, err := lsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LikeSchema ID from the query.
// Returns a *NotFoundError when no LikeSchema ID was found.
func (lsq *LikeSchemaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lsq.Limit(1).IDs(setContextOp(ctx, lsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{likeschema.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lsq *LikeSchemaQuery) FirstIDX(ctx context.Context) int {
	id, err := lsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LikeSchema entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LikeSchema entity is found.
// Returns a *NotFoundError when no LikeSchema entities are found.
func (lsq *LikeSchemaQuery) Only(ctx context.Context) (*LikeSchema, error) {
	nodes, err := lsq.Limit(2).All(setContextOp(ctx, lsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{likeschema.Label}
	default:
		return nil, &NotSingularError{likeschema.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lsq *LikeSchemaQuery) OnlyX(ctx context.Context) *LikeSchema {
	node, err := lsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LikeSchema ID in the query.
// Returns a *NotSingularError when more than one LikeSchema ID is found.
// Returns a *NotFoundError when no entities are found.
func (lsq *LikeSchemaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lsq.Limit(2).IDs(setContextOp(ctx, lsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{likeschema.Label}
	default:
		err = &NotSingularError{likeschema.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lsq *LikeSchemaQuery) OnlyIDX(ctx context.Context) int {
	id, err := lsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LikeSchemas.
func (lsq *LikeSchemaQuery) All(ctx context.Context) ([]*LikeSchema, error) {
	ctx = setContextOp(ctx, lsq.ctx, "All")
	if err := lsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LikeSchema, *LikeSchemaQuery]()
	return withInterceptors[[]*LikeSchema](ctx, lsq, qr, lsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lsq *LikeSchemaQuery) AllX(ctx context.Context) []*LikeSchema {
	nodes, err := lsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LikeSchema IDs.
func (lsq *LikeSchemaQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lsq.ctx.Unique == nil && lsq.path != nil {
		lsq.Unique(true)
	}
	ctx = setContextOp(ctx, lsq.ctx, "IDs")
	if err = lsq.Select(likeschema.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lsq *LikeSchemaQuery) IDsX(ctx context.Context) []int {
	ids, err := lsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lsq *LikeSchemaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lsq.ctx, "Count")
	if err := lsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lsq, querierCount[*LikeSchemaQuery](), lsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lsq *LikeSchemaQuery) CountX(ctx context.Context) int {
	count, err := lsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lsq *LikeSchemaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lsq.ctx, "Exist")
	switch _, err := lsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lsq *LikeSchemaQuery) ExistX(ctx context.Context) bool {
	exist, err := lsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LikeSchemaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lsq *LikeSchemaQuery) Clone() *LikeSchemaQuery {
	if lsq == nil {
		return nil
	}
	return &LikeSchemaQuery{
		config:     lsq.config,
		ctx:        lsq.ctx.Clone(),
		order:      append([]likeschema.OrderOption{}, lsq.order...),
		inters:     append([]Interceptor{}, lsq.inters...),
		predicates: append([]predicate.LikeSchema{}, lsq.predicates...),
		// clone intermediate query.
		sql:  lsq.sql.Clone(),
		path: lsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (lsq *LikeSchemaQuery) GroupBy(field string, fields ...string) *LikeSchemaGroupBy {
	lsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LikeSchemaGroupBy{build: lsq}
	grbuild.flds = &lsq.ctx.Fields
	grbuild.label = likeschema.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (lsq *LikeSchemaQuery) Select(fields ...string) *LikeSchemaSelect {
	lsq.ctx.Fields = append(lsq.ctx.Fields, fields...)
	sbuild := &LikeSchemaSelect{LikeSchemaQuery: lsq}
	sbuild.label = likeschema.Label
	sbuild.flds, sbuild.scan = &lsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LikeSchemaSelect configured with the given aggregations.
func (lsq *LikeSchemaQuery) Aggregate(fns ...AggregateFunc) *LikeSchemaSelect {
	return lsq.Select().Aggregate(fns...)
}

func (lsq *LikeSchemaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lsq); err != nil {
				return err
			}
		}
	}
	for _, f := range lsq.ctx.Fields {
		if !likeschema.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lsq.path != nil {
		prev, err := lsq.path(ctx)
		if err != nil {
			return err
		}
		lsq.sql = prev
	}
	return nil
}

func (lsq *LikeSchemaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LikeSchema, error) {
	var (
		nodes = []*LikeSchema{}
		_spec = lsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LikeSchema).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LikeSchema{config: lsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lsq *LikeSchemaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lsq.querySpec()
	_spec.Node.Columns = lsq.ctx.Fields
	if len(lsq.ctx.Fields) > 0 {
		_spec.Unique = lsq.ctx.Unique != nil && *lsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lsq.driver, _spec)
}

func (lsq *LikeSchemaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(likeschema.Table, likeschema.Columns, sqlgraph.NewFieldSpec(likeschema.FieldID, field.TypeInt))
	_spec.From = lsq.sql
	if unique := lsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lsq.path != nil {
		_spec.Unique = true
	}
	if fields := lsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, likeschema.FieldID)
		for i := range fields {
			if fields[i] != likeschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lsq *LikeSchemaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lsq.driver.Dialect())
	t1 := builder.Table(likeschema.Table)
	columns := lsq.ctx.Fields
	if len(columns) == 0 {
		columns = likeschema.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lsq.sql != nil {
		selector = lsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lsq.ctx.Unique != nil && *lsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lsq.predicates {
		p(selector)
	}
	for _, p := range lsq.order {
		p(selector)
	}
	if offset := lsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LikeSchemaGroupBy is the group-by builder for LikeSchema entities.
type LikeSchemaGroupBy struct {
	selector
	build *LikeSchemaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lsgb *LikeSchemaGroupBy) Aggregate(fns ...AggregateFunc) *LikeSchemaGroupBy {
	lsgb.fns = append(lsgb.fns, fns...)
	return lsgb
}

// Scan applies the selector query and scans the result into the given value.
func (lsgb *LikeSchemaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lsgb.build.ctx, "GroupBy")
	if err := lsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LikeSchemaQuery, *LikeSchemaGroupBy](ctx, lsgb.build, lsgb, lsgb.build.inters, v)
}

func (lsgb *LikeSchemaGroupBy) sqlScan(ctx context.Context, root *LikeSchemaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lsgb.fns))
	for _, fn := range lsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lsgb.flds)+len(lsgb.fns))
		for _, f := range *lsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LikeSchemaSelect is the builder for selecting fields of LikeSchema entities.
type LikeSchemaSelect struct {
	*LikeSchemaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lss *LikeSchemaSelect) Aggregate(fns ...AggregateFunc) *LikeSchemaSelect {
	lss.fns = append(lss.fns, fns...)
	return lss
}

// Scan applies the selector query and scans the result into the given value.
func (lss *LikeSchemaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lss.ctx, "Select")
	if err := lss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LikeSchemaQuery, *LikeSchemaSelect](ctx, lss.LikeSchemaQuery, lss, lss.inters, v)
}

func (lss *LikeSchemaSelect) sqlScan(ctx context.Context, root *LikeSchemaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lss.fns))
	for _, fn := range lss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}