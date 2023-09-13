// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"persona/libs/database/ent/logoschema"
	"persona/libs/database/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LogoSchemaUpdate is the builder for updating LogoSchema entities.
type LogoSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *LogoSchemaMutation
}

// Where appends a list predicates to the LogoSchemaUpdate builder.
func (lsu *LogoSchemaUpdate) Where(ps ...predicate.LogoSchema) *LogoSchemaUpdate {
	lsu.mutation.Where(ps...)
	return lsu
}

// Mutation returns the LogoSchemaMutation object of the builder.
func (lsu *LogoSchemaUpdate) Mutation() *LogoSchemaMutation {
	return lsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lsu *LogoSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lsu.sqlSave, lsu.mutation, lsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lsu *LogoSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := lsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lsu *LogoSchemaUpdate) Exec(ctx context.Context) error {
	_, err := lsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsu *LogoSchemaUpdate) ExecX(ctx context.Context) {
	if err := lsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lsu *LogoSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(logoschema.Table, logoschema.Columns, sqlgraph.NewFieldSpec(logoschema.FieldID, field.TypeInt))
	if ps := lsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logoschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lsu.mutation.done = true
	return n, nil
}

// LogoSchemaUpdateOne is the builder for updating a single LogoSchema entity.
type LogoSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LogoSchemaMutation
}

// Mutation returns the LogoSchemaMutation object of the builder.
func (lsuo *LogoSchemaUpdateOne) Mutation() *LogoSchemaMutation {
	return lsuo.mutation
}

// Where appends a list predicates to the LogoSchemaUpdate builder.
func (lsuo *LogoSchemaUpdateOne) Where(ps ...predicate.LogoSchema) *LogoSchemaUpdateOne {
	lsuo.mutation.Where(ps...)
	return lsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lsuo *LogoSchemaUpdateOne) Select(field string, fields ...string) *LogoSchemaUpdateOne {
	lsuo.fields = append([]string{field}, fields...)
	return lsuo
}

// Save executes the query and returns the updated LogoSchema entity.
func (lsuo *LogoSchemaUpdateOne) Save(ctx context.Context) (*LogoSchema, error) {
	return withHooks(ctx, lsuo.sqlSave, lsuo.mutation, lsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lsuo *LogoSchemaUpdateOne) SaveX(ctx context.Context) *LogoSchema {
	node, err := lsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lsuo *LogoSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := lsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsuo *LogoSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := lsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lsuo *LogoSchemaUpdateOne) sqlSave(ctx context.Context) (_node *LogoSchema, err error) {
	_spec := sqlgraph.NewUpdateSpec(logoschema.Table, logoschema.Columns, sqlgraph.NewFieldSpec(logoschema.FieldID, field.TypeInt))
	id, ok := lsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LogoSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logoschema.FieldID)
		for _, f := range fields {
			if !logoschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != logoschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &LogoSchema{config: lsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logoschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lsuo.mutation.done = true
	return _node, nil
}