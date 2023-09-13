// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"persona/libs/database/ent/phoneschema"
	"persona/libs/database/ent/predicate"
	"persona/libs/database/ent/profileschema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PhoneSchemaUpdate is the builder for updating PhoneSchema entities.
type PhoneSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *PhoneSchemaMutation
}

// Where appends a list predicates to the PhoneSchemaUpdate builder.
func (psu *PhoneSchemaUpdate) Where(ps ...predicate.PhoneSchema) *PhoneSchemaUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetOwnerID sets the "owner" edge to the ProfileSchema entity by ID.
func (psu *PhoneSchemaUpdate) SetOwnerID(id int) *PhoneSchemaUpdate {
	psu.mutation.SetOwnerID(id)
	return psu
}

// SetOwner sets the "owner" edge to the ProfileSchema entity.
func (psu *PhoneSchemaUpdate) SetOwner(p *ProfileSchema) *PhoneSchemaUpdate {
	return psu.SetOwnerID(p.ID)
}

// Mutation returns the PhoneSchemaMutation object of the builder.
func (psu *PhoneSchemaUpdate) Mutation() *PhoneSchemaMutation {
	return psu.mutation
}

// ClearOwner clears the "owner" edge to the ProfileSchema entity.
func (psu *PhoneSchemaUpdate) ClearOwner() *PhoneSchemaUpdate {
	psu.mutation.ClearOwner()
	return psu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PhoneSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PhoneSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PhoneSchemaUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PhoneSchemaUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psu *PhoneSchemaUpdate) check() error {
	if _, ok := psu.mutation.OwnerID(); psu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PhoneSchema.owner"`)
	}
	return nil
}

func (psu *PhoneSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := psu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(phoneschema.Table, phoneschema.Columns, sqlgraph.NewFieldSpec(phoneschema.FieldID, field.TypeInt))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if psu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   phoneschema.OwnerTable,
			Columns: []string{phoneschema.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   phoneschema.OwnerTable,
			Columns: []string{phoneschema.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phoneschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// PhoneSchemaUpdateOne is the builder for updating a single PhoneSchema entity.
type PhoneSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PhoneSchemaMutation
}

// SetOwnerID sets the "owner" edge to the ProfileSchema entity by ID.
func (psuo *PhoneSchemaUpdateOne) SetOwnerID(id int) *PhoneSchemaUpdateOne {
	psuo.mutation.SetOwnerID(id)
	return psuo
}

// SetOwner sets the "owner" edge to the ProfileSchema entity.
func (psuo *PhoneSchemaUpdateOne) SetOwner(p *ProfileSchema) *PhoneSchemaUpdateOne {
	return psuo.SetOwnerID(p.ID)
}

// Mutation returns the PhoneSchemaMutation object of the builder.
func (psuo *PhoneSchemaUpdateOne) Mutation() *PhoneSchemaMutation {
	return psuo.mutation
}

// ClearOwner clears the "owner" edge to the ProfileSchema entity.
func (psuo *PhoneSchemaUpdateOne) ClearOwner() *PhoneSchemaUpdateOne {
	psuo.mutation.ClearOwner()
	return psuo
}

// Where appends a list predicates to the PhoneSchemaUpdate builder.
func (psuo *PhoneSchemaUpdateOne) Where(ps ...predicate.PhoneSchema) *PhoneSchemaUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PhoneSchemaUpdateOne) Select(field string, fields ...string) *PhoneSchemaUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PhoneSchema entity.
func (psuo *PhoneSchemaUpdateOne) Save(ctx context.Context) (*PhoneSchema, error) {
	return withHooks(ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PhoneSchemaUpdateOne) SaveX(ctx context.Context) *PhoneSchema {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PhoneSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PhoneSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psuo *PhoneSchemaUpdateOne) check() error {
	if _, ok := psuo.mutation.OwnerID(); psuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PhoneSchema.owner"`)
	}
	return nil
}

func (psuo *PhoneSchemaUpdateOne) sqlSave(ctx context.Context) (_node *PhoneSchema, err error) {
	if err := psuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(phoneschema.Table, phoneschema.Columns, sqlgraph.NewFieldSpec(phoneschema.FieldID, field.TypeInt))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PhoneSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, phoneschema.FieldID)
		for _, f := range fields {
			if !phoneschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != phoneschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if psuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   phoneschema.OwnerTable,
			Columns: []string{phoneschema.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   phoneschema.OwnerTable,
			Columns: []string{phoneschema.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(profileschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PhoneSchema{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{phoneschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}
