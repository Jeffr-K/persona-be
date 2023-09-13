// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"persona/libs/database/ent/paymentschema"
	"persona/libs/database/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PaymentSchemaUpdate is the builder for updating PaymentSchema entities.
type PaymentSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *PaymentSchemaMutation
}

// Where appends a list predicates to the PaymentSchemaUpdate builder.
func (psu *PaymentSchemaUpdate) Where(ps ...predicate.PaymentSchema) *PaymentSchemaUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// Mutation returns the PaymentSchemaMutation object of the builder.
func (psu *PaymentSchemaUpdate) Mutation() *PaymentSchemaMutation {
	return psu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PaymentSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PaymentSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PaymentSchemaUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PaymentSchemaUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psu *PaymentSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(paymentschema.Table, paymentschema.Columns, sqlgraph.NewFieldSpec(paymentschema.FieldID, field.TypeInt))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// PaymentSchemaUpdateOne is the builder for updating a single PaymentSchema entity.
type PaymentSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PaymentSchemaMutation
}

// Mutation returns the PaymentSchemaMutation object of the builder.
func (psuo *PaymentSchemaUpdateOne) Mutation() *PaymentSchemaMutation {
	return psuo.mutation
}

// Where appends a list predicates to the PaymentSchemaUpdate builder.
func (psuo *PaymentSchemaUpdateOne) Where(ps ...predicate.PaymentSchema) *PaymentSchemaUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PaymentSchemaUpdateOne) Select(field string, fields ...string) *PaymentSchemaUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PaymentSchema entity.
func (psuo *PaymentSchemaUpdateOne) Save(ctx context.Context) (*PaymentSchema, error) {
	return withHooks(ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PaymentSchemaUpdateOne) SaveX(ctx context.Context) *PaymentSchema {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PaymentSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PaymentSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (psuo *PaymentSchemaUpdateOne) sqlSave(ctx context.Context) (_node *PaymentSchema, err error) {
	_spec := sqlgraph.NewUpdateSpec(paymentschema.Table, paymentschema.Columns, sqlgraph.NewFieldSpec(paymentschema.FieldID, field.TypeInt))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PaymentSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymentschema.FieldID)
		for _, f := range fields {
			if !paymentschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != paymentschema.FieldID {
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
	_node = &PaymentSchema{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}