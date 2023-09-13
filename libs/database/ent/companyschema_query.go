// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"persona/libs/database/ent/companyschema"
	"persona/libs/database/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanySchemaQuery is the builder for querying CompanySchema entities.
type CompanySchemaQuery struct {
	config
	ctx        *QueryContext
	order      []companyschema.OrderOption
	inters     []Interceptor
	predicates []predicate.CompanySchema
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CompanySchemaQuery builder.
func (csq *CompanySchemaQuery) Where(ps ...predicate.CompanySchema) *CompanySchemaQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit the number of records to be returned by this query.
func (csq *CompanySchemaQuery) Limit(limit int) *CompanySchemaQuery {
	csq.ctx.Limit = &limit
	return csq
}

// Offset to start from.
func (csq *CompanySchemaQuery) Offset(offset int) *CompanySchemaQuery {
	csq.ctx.Offset = &offset
	return csq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csq *CompanySchemaQuery) Unique(unique bool) *CompanySchemaQuery {
	csq.ctx.Unique = &unique
	return csq
}

// Order specifies how the records should be ordered.
func (csq *CompanySchemaQuery) Order(o ...companyschema.OrderOption) *CompanySchemaQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// First returns the first CompanySchema entity from the query.
// Returns a *NotFoundError when no CompanySchema was found.
func (csq *CompanySchemaQuery) First(ctx context.Context) (*CompanySchema, error) {
	nodes, err := csq.Limit(1).All(setContextOp(ctx, csq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{companyschema.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *CompanySchemaQuery) FirstX(ctx context.Context) *CompanySchema {
	node, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CompanySchema ID from the query.
// Returns a *NotFoundError when no CompanySchema ID was found.
func (csq *CompanySchemaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(1).IDs(setContextOp(ctx, csq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{companyschema.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csq *CompanySchemaQuery) FirstIDX(ctx context.Context) int {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CompanySchema entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CompanySchema entity is found.
// Returns a *NotFoundError when no CompanySchema entities are found.
func (csq *CompanySchemaQuery) Only(ctx context.Context) (*CompanySchema, error) {
	nodes, err := csq.Limit(2).All(setContextOp(ctx, csq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{companyschema.Label}
	default:
		return nil, &NotSingularError{companyschema.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *CompanySchemaQuery) OnlyX(ctx context.Context) *CompanySchema {
	node, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CompanySchema ID in the query.
// Returns a *NotSingularError when more than one CompanySchema ID is found.
// Returns a *NotFoundError when no entities are found.
func (csq *CompanySchemaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(2).IDs(setContextOp(ctx, csq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{companyschema.Label}
	default:
		err = &NotSingularError{companyschema.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *CompanySchemaQuery) OnlyIDX(ctx context.Context) int {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CompanySchemas.
func (csq *CompanySchemaQuery) All(ctx context.Context) ([]*CompanySchema, error) {
	ctx = setContextOp(ctx, csq.ctx, "All")
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CompanySchema, *CompanySchemaQuery]()
	return withInterceptors[[]*CompanySchema](ctx, csq, qr, csq.inters)
}

// AllX is like All, but panics if an error occurs.
func (csq *CompanySchemaQuery) AllX(ctx context.Context) []*CompanySchema {
	nodes, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CompanySchema IDs.
func (csq *CompanySchemaQuery) IDs(ctx context.Context) (ids []int, err error) {
	if csq.ctx.Unique == nil && csq.path != nil {
		csq.Unique(true)
	}
	ctx = setContextOp(ctx, csq.ctx, "IDs")
	if err = csq.Select(companyschema.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *CompanySchemaQuery) IDsX(ctx context.Context) []int {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *CompanySchemaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, csq.ctx, "Count")
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, csq, querierCount[*CompanySchemaQuery](), csq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (csq *CompanySchemaQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *CompanySchemaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, csq.ctx, "Exist")
	switch _, err := csq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *CompanySchemaQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CompanySchemaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *CompanySchemaQuery) Clone() *CompanySchemaQuery {
	if csq == nil {
		return nil
	}
	return &CompanySchemaQuery{
		config:     csq.config,
		ctx:        csq.ctx.Clone(),
		order:      append([]companyschema.OrderOption{}, csq.order...),
		inters:     append([]Interceptor{}, csq.inters...),
		predicates: append([]predicate.CompanySchema{}, csq.predicates...),
		// clone intermediate query.
		sql:  csq.sql.Clone(),
		path: csq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (csq *CompanySchemaQuery) GroupBy(field string, fields ...string) *CompanySchemaGroupBy {
	csq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CompanySchemaGroupBy{build: csq}
	grbuild.flds = &csq.ctx.Fields
	grbuild.label = companyschema.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (csq *CompanySchemaQuery) Select(fields ...string) *CompanySchemaSelect {
	csq.ctx.Fields = append(csq.ctx.Fields, fields...)
	sbuild := &CompanySchemaSelect{CompanySchemaQuery: csq}
	sbuild.label = companyschema.Label
	sbuild.flds, sbuild.scan = &csq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CompanySchemaSelect configured with the given aggregations.
func (csq *CompanySchemaQuery) Aggregate(fns ...AggregateFunc) *CompanySchemaSelect {
	return csq.Select().Aggregate(fns...)
}

func (csq *CompanySchemaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range csq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, csq); err != nil {
				return err
			}
		}
	}
	for _, f := range csq.ctx.Fields {
		if !companyschema.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *CompanySchemaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CompanySchema, error) {
	var (
		nodes = []*CompanySchema{}
		_spec = csq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CompanySchema).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CompanySchema{config: csq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (csq *CompanySchemaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	_spec.Node.Columns = csq.ctx.Fields
	if len(csq.ctx.Fields) > 0 {
		_spec.Unique = csq.ctx.Unique != nil && *csq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *CompanySchemaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(companyschema.Table, companyschema.Columns, sqlgraph.NewFieldSpec(companyschema.FieldID, field.TypeInt))
	_spec.From = csq.sql
	if unique := csq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if csq.path != nil {
		_spec.Unique = true
	}
	if fields := csq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, companyschema.FieldID)
		for i := range fields {
			if fields[i] != companyschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csq *CompanySchemaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(companyschema.Table)
	columns := csq.ctx.Fields
	if len(columns) == 0 {
		columns = companyschema.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csq.ctx.Unique != nil && *csq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CompanySchemaGroupBy is the group-by builder for CompanySchema entities.
type CompanySchemaGroupBy struct {
	selector
	build *CompanySchemaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *CompanySchemaGroupBy) Aggregate(fns ...AggregateFunc) *CompanySchemaGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the selector query and scans the result into the given value.
func (csgb *CompanySchemaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csgb.build.ctx, "GroupBy")
	if err := csgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanySchemaQuery, *CompanySchemaGroupBy](ctx, csgb.build, csgb, csgb.build.inters, v)
}

func (csgb *CompanySchemaGroupBy) sqlScan(ctx context.Context, root *CompanySchemaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(csgb.fns))
	for _, fn := range csgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*csgb.flds)+len(csgb.fns))
		for _, f := range *csgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*csgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CompanySchemaSelect is the builder for selecting fields of CompanySchema entities.
type CompanySchemaSelect struct {
	*CompanySchemaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (css *CompanySchemaSelect) Aggregate(fns ...AggregateFunc) *CompanySchemaSelect {
	css.fns = append(css.fns, fns...)
	return css
}

// Scan applies the selector query and scans the result into the given value.
func (css *CompanySchemaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, css.ctx, "Select")
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CompanySchemaQuery, *CompanySchemaSelect](ctx, css.CompanySchemaQuery, css, css.inters, v)
}

func (css *CompanySchemaSelect) sqlScan(ctx context.Context, root *CompanySchemaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(css.fns))
	for _, fn := range css.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*css.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}