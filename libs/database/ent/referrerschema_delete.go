// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"persona/libs/database/ent/predicate"
	"persona/libs/database/ent/referrerschema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReferrerSchemaDelete is the builder for deleting a ReferrerSchema entity.
type ReferrerSchemaDelete struct {
	config
	hooks    []Hook
	mutation *ReferrerSchemaMutation
}

// Where appends a list predicates to the ReferrerSchemaDelete builder.
func (rsd *ReferrerSchemaDelete) Where(ps ...predicate.ReferrerSchema) *ReferrerSchemaDelete {
	rsd.mutation.Where(ps...)
	return rsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rsd *ReferrerSchemaDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rsd.sqlExec, rsd.mutation, rsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rsd *ReferrerSchemaDelete) ExecX(ctx context.Context) int {
	n, err := rsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rsd *ReferrerSchemaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(referrerschema.Table, sqlgraph.NewFieldSpec(referrerschema.FieldID, field.TypeInt))
	if ps := rsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rsd.mutation.done = true
	return affected, err
}

// ReferrerSchemaDeleteOne is the builder for deleting a single ReferrerSchema entity.
type ReferrerSchemaDeleteOne struct {
	rsd *ReferrerSchemaDelete
}

// Where appends a list predicates to the ReferrerSchemaDelete builder.
func (rsdo *ReferrerSchemaDeleteOne) Where(ps ...predicate.ReferrerSchema) *ReferrerSchemaDeleteOne {
	rsdo.rsd.mutation.Where(ps...)
	return rsdo
}

// Exec executes the deletion query.
func (rsdo *ReferrerSchemaDeleteOne) Exec(ctx context.Context) error {
	n, err := rsdo.rsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{referrerschema.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rsdo *ReferrerSchemaDeleteOne) ExecX(ctx context.Context) {
	if err := rsdo.Exec(ctx); err != nil {
		panic(err)
	}
}