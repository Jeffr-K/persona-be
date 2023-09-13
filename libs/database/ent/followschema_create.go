// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"persona/libs/database/ent/followschema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FollowSchemaCreate is the builder for creating a FollowSchema entity.
type FollowSchemaCreate struct {
	config
	mutation *FollowSchemaMutation
	hooks    []Hook
}

// Mutation returns the FollowSchemaMutation object of the builder.
func (fsc *FollowSchemaCreate) Mutation() *FollowSchemaMutation {
	return fsc.mutation
}

// Save creates the FollowSchema in the database.
func (fsc *FollowSchemaCreate) Save(ctx context.Context) (*FollowSchema, error) {
	return withHooks(ctx, fsc.sqlSave, fsc.mutation, fsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fsc *FollowSchemaCreate) SaveX(ctx context.Context) *FollowSchema {
	v, err := fsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fsc *FollowSchemaCreate) Exec(ctx context.Context) error {
	_, err := fsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsc *FollowSchemaCreate) ExecX(ctx context.Context) {
	if err := fsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fsc *FollowSchemaCreate) check() error {
	return nil
}

func (fsc *FollowSchemaCreate) sqlSave(ctx context.Context) (*FollowSchema, error) {
	if err := fsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fsc.mutation.id = &_node.ID
	fsc.mutation.done = true
	return _node, nil
}

func (fsc *FollowSchemaCreate) createSpec() (*FollowSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &FollowSchema{config: fsc.config}
		_spec = sqlgraph.NewCreateSpec(followschema.Table, sqlgraph.NewFieldSpec(followschema.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// FollowSchemaCreateBulk is the builder for creating many FollowSchema entities in bulk.
type FollowSchemaCreateBulk struct {
	config
	builders []*FollowSchemaCreate
}

// Save creates the FollowSchema entities in the database.
func (fscb *FollowSchemaCreateBulk) Save(ctx context.Context) ([]*FollowSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fscb.builders))
	nodes := make([]*FollowSchema, len(fscb.builders))
	mutators := make([]Mutator, len(fscb.builders))
	for i := range fscb.builders {
		func(i int, root context.Context) {
			builder := fscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FollowSchemaMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fscb *FollowSchemaCreateBulk) SaveX(ctx context.Context) []*FollowSchema {
	v, err := fscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fscb *FollowSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := fscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fscb *FollowSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := fscb.Exec(ctx); err != nil {
		panic(err)
	}
}