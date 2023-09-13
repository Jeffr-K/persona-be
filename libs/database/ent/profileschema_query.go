// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"persona/libs/database/ent/phoneschema"
	"persona/libs/database/ent/photoschema"
	"persona/libs/database/ent/predicate"
	"persona/libs/database/ent/profileschema"
	"persona/libs/database/ent/userschema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileSchemaQuery is the builder for querying ProfileSchema entities.
type ProfileSchemaQuery struct {
	config
	ctx        *QueryContext
	order      []profileschema.OrderOption
	inters     []Interceptor
	predicates []predicate.ProfileSchema
	withOwner  *UserSchemaQuery
	withPhoto  *PhotoSchemaQuery
	withPhone  *PhoneSchemaQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProfileSchemaQuery builder.
func (psq *ProfileSchemaQuery) Where(ps ...predicate.ProfileSchema) *ProfileSchemaQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit the number of records to be returned by this query.
func (psq *ProfileSchemaQuery) Limit(limit int) *ProfileSchemaQuery {
	psq.ctx.Limit = &limit
	return psq
}

// Offset to start from.
func (psq *ProfileSchemaQuery) Offset(offset int) *ProfileSchemaQuery {
	psq.ctx.Offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *ProfileSchemaQuery) Unique(unique bool) *ProfileSchemaQuery {
	psq.ctx.Unique = &unique
	return psq
}

// Order specifies how the records should be ordered.
func (psq *ProfileSchemaQuery) Order(o ...profileschema.OrderOption) *ProfileSchemaQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// QueryOwner chains the current query on the "owner" edge.
func (psq *ProfileSchemaQuery) QueryOwner() *UserSchemaQuery {
	query := (&UserSchemaClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profileschema.Table, profileschema.FieldID, selector),
			sqlgraph.To(userschema.Table, userschema.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, profileschema.OwnerTable, profileschema.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPhoto chains the current query on the "photo" edge.
func (psq *ProfileSchemaQuery) QueryPhoto() *PhotoSchemaQuery {
	query := (&PhotoSchemaClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profileschema.Table, profileschema.FieldID, selector),
			sqlgraph.To(photoschema.Table, photoschema.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, profileschema.PhotoTable, profileschema.PhotoColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPhone chains the current query on the "phone" edge.
func (psq *ProfileSchemaQuery) QueryPhone() *PhoneSchemaQuery {
	query := (&PhoneSchemaClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(profileschema.Table, profileschema.FieldID, selector),
			sqlgraph.To(phoneschema.Table, phoneschema.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, profileschema.PhoneTable, profileschema.PhoneColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProfileSchema entity from the query.
// Returns a *NotFoundError when no ProfileSchema was found.
func (psq *ProfileSchemaQuery) First(ctx context.Context) (*ProfileSchema, error) {
	nodes, err := psq.Limit(1).All(setContextOp(ctx, psq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{profileschema.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *ProfileSchemaQuery) FirstX(ctx context.Context) *ProfileSchema {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProfileSchema ID from the query.
// Returns a *NotFoundError when no ProfileSchema ID was found.
func (psq *ProfileSchemaQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(1).IDs(setContextOp(ctx, psq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{profileschema.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *ProfileSchemaQuery) FirstIDX(ctx context.Context) int {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProfileSchema entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProfileSchema entity is found.
// Returns a *NotFoundError when no ProfileSchema entities are found.
func (psq *ProfileSchemaQuery) Only(ctx context.Context) (*ProfileSchema, error) {
	nodes, err := psq.Limit(2).All(setContextOp(ctx, psq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{profileschema.Label}
	default:
		return nil, &NotSingularError{profileschema.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *ProfileSchemaQuery) OnlyX(ctx context.Context) *ProfileSchema {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProfileSchema ID in the query.
// Returns a *NotSingularError when more than one ProfileSchema ID is found.
// Returns a *NotFoundError when no entities are found.
func (psq *ProfileSchemaQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(2).IDs(setContextOp(ctx, psq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{profileschema.Label}
	default:
		err = &NotSingularError{profileschema.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *ProfileSchemaQuery) OnlyIDX(ctx context.Context) int {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProfileSchemas.
func (psq *ProfileSchemaQuery) All(ctx context.Context) ([]*ProfileSchema, error) {
	ctx = setContextOp(ctx, psq.ctx, "All")
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProfileSchema, *ProfileSchemaQuery]()
	return withInterceptors[[]*ProfileSchema](ctx, psq, qr, psq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psq *ProfileSchemaQuery) AllX(ctx context.Context) []*ProfileSchema {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProfileSchema IDs.
func (psq *ProfileSchemaQuery) IDs(ctx context.Context) (ids []int, err error) {
	if psq.ctx.Unique == nil && psq.path != nil {
		psq.Unique(true)
	}
	ctx = setContextOp(ctx, psq.ctx, "IDs")
	if err = psq.Select(profileschema.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *ProfileSchemaQuery) IDsX(ctx context.Context) []int {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *ProfileSchemaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psq.ctx, "Count")
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psq, querierCount[*ProfileSchemaQuery](), psq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psq *ProfileSchemaQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *ProfileSchemaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psq.ctx, "Exist")
	switch _, err := psq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *ProfileSchemaQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProfileSchemaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *ProfileSchemaQuery) Clone() *ProfileSchemaQuery {
	if psq == nil {
		return nil
	}
	return &ProfileSchemaQuery{
		config:     psq.config,
		ctx:        psq.ctx.Clone(),
		order:      append([]profileschema.OrderOption{}, psq.order...),
		inters:     append([]Interceptor{}, psq.inters...),
		predicates: append([]predicate.ProfileSchema{}, psq.predicates...),
		withOwner:  psq.withOwner.Clone(),
		withPhoto:  psq.withPhoto.Clone(),
		withPhone:  psq.withPhone.Clone(),
		// clone intermediate query.
		sql:  psq.sql.Clone(),
		path: psq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *ProfileSchemaQuery) WithOwner(opts ...func(*UserSchemaQuery)) *ProfileSchemaQuery {
	query := (&UserSchemaClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withOwner = query
	return psq
}

// WithPhoto tells the query-builder to eager-load the nodes that are connected to
// the "photo" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *ProfileSchemaQuery) WithPhoto(opts ...func(*PhotoSchemaQuery)) *ProfileSchemaQuery {
	query := (&PhotoSchemaClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withPhoto = query
	return psq
}

// WithPhone tells the query-builder to eager-load the nodes that are connected to
// the "phone" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *ProfileSchemaQuery) WithPhone(opts ...func(*PhoneSchemaQuery)) *ProfileSchemaQuery {
	query := (&PhoneSchemaClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withPhone = query
	return psq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (psq *ProfileSchemaQuery) GroupBy(field string, fields ...string) *ProfileSchemaGroupBy {
	psq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProfileSchemaGroupBy{build: psq}
	grbuild.flds = &psq.ctx.Fields
	grbuild.label = profileschema.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (psq *ProfileSchemaQuery) Select(fields ...string) *ProfileSchemaSelect {
	psq.ctx.Fields = append(psq.ctx.Fields, fields...)
	sbuild := &ProfileSchemaSelect{ProfileSchemaQuery: psq}
	sbuild.label = profileschema.Label
	sbuild.flds, sbuild.scan = &psq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProfileSchemaSelect configured with the given aggregations.
func (psq *ProfileSchemaQuery) Aggregate(fns ...AggregateFunc) *ProfileSchemaSelect {
	return psq.Select().Aggregate(fns...)
}

func (psq *ProfileSchemaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psq); err != nil {
				return err
			}
		}
	}
	for _, f := range psq.ctx.Fields {
		if !profileschema.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *ProfileSchemaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProfileSchema, error) {
	var (
		nodes       = []*ProfileSchema{}
		withFKs     = psq.withFKs
		_spec       = psq.querySpec()
		loadedTypes = [3]bool{
			psq.withOwner != nil,
			psq.withPhoto != nil,
			psq.withPhone != nil,
		}
	)
	if psq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, profileschema.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProfileSchema).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProfileSchema{config: psq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := psq.withOwner; query != nil {
		if err := psq.loadOwner(ctx, query, nodes, nil,
			func(n *ProfileSchema, e *UserSchema) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := psq.withPhoto; query != nil {
		if err := psq.loadPhoto(ctx, query, nodes, nil,
			func(n *ProfileSchema, e *PhotoSchema) { n.Edges.Photo = e }); err != nil {
			return nil, err
		}
	}
	if query := psq.withPhone; query != nil {
		if err := psq.loadPhone(ctx, query, nodes, nil,
			func(n *ProfileSchema, e *PhoneSchema) { n.Edges.Phone = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (psq *ProfileSchemaQuery) loadOwner(ctx context.Context, query *UserSchemaQuery, nodes []*ProfileSchema, init func(*ProfileSchema), assign func(*ProfileSchema, *UserSchema)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProfileSchema)
	for i := range nodes {
		if nodes[i].user_schema_profile == nil {
			continue
		}
		fk := *nodes[i].user_schema_profile
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(userschema.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_schema_profile" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (psq *ProfileSchemaQuery) loadPhoto(ctx context.Context, query *PhotoSchemaQuery, nodes []*ProfileSchema, init func(*ProfileSchema), assign func(*ProfileSchema, *PhotoSchema)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ProfileSchema)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.PhotoSchema(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(profileschema.PhotoColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.profile_schema_photo
		if fk == nil {
			return fmt.Errorf(`foreign-key "profile_schema_photo" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "profile_schema_photo" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (psq *ProfileSchemaQuery) loadPhone(ctx context.Context, query *PhoneSchemaQuery, nodes []*ProfileSchema, init func(*ProfileSchema), assign func(*ProfileSchema, *PhoneSchema)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ProfileSchema)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.PhoneSchema(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(profileschema.PhoneColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.profile_schema_phone
		if fk == nil {
			return fmt.Errorf(`foreign-key "profile_schema_phone" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "profile_schema_phone" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (psq *ProfileSchemaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	_spec.Node.Columns = psq.ctx.Fields
	if len(psq.ctx.Fields) > 0 {
		_spec.Unique = psq.ctx.Unique != nil && *psq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *ProfileSchemaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(profileschema.Table, profileschema.Columns, sqlgraph.NewFieldSpec(profileschema.FieldID, field.TypeInt))
	_spec.From = psq.sql
	if unique := psq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psq.path != nil {
		_spec.Unique = true
	}
	if fields := psq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profileschema.FieldID)
		for i := range fields {
			if fields[i] != profileschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *ProfileSchemaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(profileschema.Table)
	columns := psq.ctx.Fields
	if len(columns) == 0 {
		columns = profileschema.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psq.ctx.Unique != nil && *psq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProfileSchemaGroupBy is the group-by builder for ProfileSchema entities.
type ProfileSchemaGroupBy struct {
	selector
	build *ProfileSchemaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *ProfileSchemaGroupBy) Aggregate(fns ...AggregateFunc) *ProfileSchemaGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the selector query and scans the result into the given value.
func (psgb *ProfileSchemaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psgb.build.ctx, "GroupBy")
	if err := psgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfileSchemaQuery, *ProfileSchemaGroupBy](ctx, psgb.build, psgb, psgb.build.inters, v)
}

func (psgb *ProfileSchemaGroupBy) sqlScan(ctx context.Context, root *ProfileSchemaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psgb.flds)+len(psgb.fns))
		for _, f := range *psgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProfileSchemaSelect is the builder for selecting fields of ProfileSchema entities.
type ProfileSchemaSelect struct {
	*ProfileSchemaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pss *ProfileSchemaSelect) Aggregate(fns ...AggregateFunc) *ProfileSchemaSelect {
	pss.fns = append(pss.fns, fns...)
	return pss
}

// Scan applies the selector query and scans the result into the given value.
func (pss *ProfileSchemaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pss.ctx, "Select")
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfileSchemaQuery, *ProfileSchemaSelect](ctx, pss.ProfileSchemaQuery, pss, pss.inters, v)
}

func (pss *ProfileSchemaSelect) sqlScan(ctx context.Context, root *ProfileSchemaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pss.fns))
	for _, fn := range pss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
