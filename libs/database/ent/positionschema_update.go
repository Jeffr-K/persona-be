// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"persona/libs/database/ent/positionschema"
	"persona/libs/database/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PositionSchemaUpdate is the builder for updating PositionSchema entities.
type PositionSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *PositionSchemaMutation
}

// Where appends a list predicates to the PositionSchemaUpdate builder.
func (psu *PositionSchemaUpdate) Where(ps ...predicate.PositionSchema) *PositionSchemaUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// Mutation returns the PositionSchemaMutation object of the builder.
func (psu *PositionSchemaUpdate) Mutation() *PositionSchemaMutation {
	return psu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PositionSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PositionSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PositionSchemaUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PositionSchemaUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psu *PositionSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(positionschema.Table, positionschema.Columns, sqlgraph.NewFieldSpec(positionschema.FieldID, field.TypeInt))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{positionschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// PositionSchemaUpdateOne is the builder for updating a single PositionSchema entity.
type PositionSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PositionSchemaMutation
}

// Mutation returns the PositionSchemaMutation object of the builder.
func (psuo *PositionSchemaUpdateOne) Mutation() *PositionSchemaMutation {
	return psuo.mutation
}

// Where appends a list predicates to the PositionSchemaUpdate builder.
func (psuo *PositionSchemaUpdateOne) Where(ps ...predicate.PositionSchema) *PositionSchemaUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PositionSchemaUpdateOne) Select(field string, fields ...string) *PositionSchemaUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PositionSchema entity.
func (psuo *PositionSchemaUpdateOne) Save(ctx context.Context) (*PositionSchema, error) {
	return withHooks(ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PositionSchemaUpdateOne) SaveX(ctx context.Context) *PositionSchema {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PositionSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PositionSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psuo *PositionSchemaUpdateOne) sqlSave(ctx context.Context) (_node *PositionSchema, err error) {
	_spec := sqlgraph.NewUpdateSpec(positionschema.Table, positionschema.Columns, sqlgraph.NewFieldSpec(positionschema.FieldID, field.TypeInt))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PositionSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, positionschema.FieldID)
		for _, f := range fields {
			if !positionschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != positionschema.FieldID {
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
	_node = &PositionSchema{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{positionschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}