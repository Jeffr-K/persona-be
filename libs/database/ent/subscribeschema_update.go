// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"persona/libs/database/ent/predicate"
	"persona/libs/database/ent/subscribeschema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubscribeSchemaUpdate is the builder for updating SubscribeSchema entities.
type SubscribeSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *SubscribeSchemaMutation
}

// Where appends a list predicates to the SubscribeSchemaUpdate builder.
func (ssu *SubscribeSchemaUpdate) Where(ps ...predicate.SubscribeSchema) *SubscribeSchemaUpdate {
	ssu.mutation.Where(ps...)
	return ssu
}

// Mutation returns the SubscribeSchemaMutation object of the builder.
func (ssu *SubscribeSchemaUpdate) Mutation() *SubscribeSchemaMutation {
	return ssu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ssu *SubscribeSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ssu.sqlSave, ssu.mutation, ssu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ssu *SubscribeSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := ssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ssu *SubscribeSchemaUpdate) Exec(ctx context.Context) error {
	_, err := ssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssu *SubscribeSchemaUpdate) ExecX(ctx context.Context) {
	if err := ssu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssu *SubscribeSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscribeschema.Table, subscribeschema.Columns, sqlgraph.NewFieldSpec(subscribeschema.FieldID, field.TypeInt))
	if ps := ssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscribeschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ssu.mutation.done = true
	return n, nil
}

// SubscribeSchemaUpdateOne is the builder for updating a single SubscribeSchema entity.
type SubscribeSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscribeSchemaMutation
}

// Mutation returns the SubscribeSchemaMutation object of the builder.
func (ssuo *SubscribeSchemaUpdateOne) Mutation() *SubscribeSchemaMutation {
	return ssuo.mutation
}

// Where appends a list predicates to the SubscribeSchemaUpdate builder.
func (ssuo *SubscribeSchemaUpdateOne) Where(ps ...predicate.SubscribeSchema) *SubscribeSchemaUpdateOne {
	ssuo.mutation.Where(ps...)
	return ssuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ssuo *SubscribeSchemaUpdateOne) Select(field string, fields ...string) *SubscribeSchemaUpdateOne {
	ssuo.fields = append([]string{field}, fields...)
	return ssuo
}

// Save executes the query and returns the updated SubscribeSchema entity.
func (ssuo *SubscribeSchemaUpdateOne) Save(ctx context.Context) (*SubscribeSchema, error) {
	return withHooks(ctx, ssuo.sqlSave, ssuo.mutation, ssuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ssuo *SubscribeSchemaUpdateOne) SaveX(ctx context.Context) *SubscribeSchema {
	node, err := ssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ssuo *SubscribeSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := ssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssuo *SubscribeSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := ssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssuo *SubscribeSchemaUpdateOne) sqlSave(ctx context.Context) (_node *SubscribeSchema, err error) {
	_spec := sqlgraph.NewUpdateSpec(subscribeschema.Table, subscribeschema.Columns, sqlgraph.NewFieldSpec(subscribeschema.FieldID, field.TypeInt))
	id, ok := ssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SubscribeSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscribeschema.FieldID)
		for _, f := range fields {
			if !subscribeschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscribeschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ssuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &SubscribeSchema{config: ssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscribeschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ssuo.mutation.done = true
	return _node, nil
}
