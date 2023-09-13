// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"persona/libs/database/ent/predicate"
	"persona/libs/database/ent/skillschema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SkillSchemaUpdate is the builder for updating SkillSchema entities.
type SkillSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *SkillSchemaMutation
}

// Where appends a list predicates to the SkillSchemaUpdate builder.
func (ssu *SkillSchemaUpdate) Where(ps ...predicate.SkillSchema) *SkillSchemaUpdate {
	ssu.mutation.Where(ps...)
	return ssu
}

// Mutation returns the SkillSchemaMutation object of the builder.
func (ssu *SkillSchemaUpdate) Mutation() *SkillSchemaMutation {
	return ssu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ssu *SkillSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ssu.sqlSave, ssu.mutation, ssu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ssu *SkillSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := ssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ssu *SkillSchemaUpdate) Exec(ctx context.Context) error {
	_, err := ssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssu *SkillSchemaUpdate) ExecX(ctx context.Context) {
	if err := ssu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssu *SkillSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(skillschema.Table, skillschema.Columns, sqlgraph.NewFieldSpec(skillschema.FieldID, field.TypeInt))
	if ps := ssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skillschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ssu.mutation.done = true
	return n, nil
}

// SkillSchemaUpdateOne is the builder for updating a single SkillSchema entity.
type SkillSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkillSchemaMutation
}

// Mutation returns the SkillSchemaMutation object of the builder.
func (ssuo *SkillSchemaUpdateOne) Mutation() *SkillSchemaMutation {
	return ssuo.mutation
}

// Where appends a list predicates to the SkillSchemaUpdate builder.
func (ssuo *SkillSchemaUpdateOne) Where(ps ...predicate.SkillSchema) *SkillSchemaUpdateOne {
	ssuo.mutation.Where(ps...)
	return ssuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ssuo *SkillSchemaUpdateOne) Select(field string, fields ...string) *SkillSchemaUpdateOne {
	ssuo.fields = append([]string{field}, fields...)
	return ssuo
}

// Save executes the query and returns the updated SkillSchema entity.
func (ssuo *SkillSchemaUpdateOne) Save(ctx context.Context) (*SkillSchema, error) {
	return withHooks(ctx, ssuo.sqlSave, ssuo.mutation, ssuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ssuo *SkillSchemaUpdateOne) SaveX(ctx context.Context) *SkillSchema {
	node, err := ssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ssuo *SkillSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := ssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssuo *SkillSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := ssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssuo *SkillSchemaUpdateOne) sqlSave(ctx context.Context) (_node *SkillSchema, err error) {
	_spec := sqlgraph.NewUpdateSpec(skillschema.Table, skillschema.Columns, sqlgraph.NewFieldSpec(skillschema.FieldID, field.TypeInt))
	id, ok := ssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SkillSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skillschema.FieldID)
		for _, f := range fields {
			if !skillschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skillschema.FieldID {
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
	_node = &SkillSchema{config: ssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skillschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ssuo.mutation.done = true
	return _node, nil
}