// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"persona/libs/database/ent/threadschema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThreadSchemaCreate is the builder for creating a ThreadSchema entity.
type ThreadSchemaCreate struct {
	config
	mutation *ThreadSchemaMutation
	hooks    []Hook
}

// Mutation returns the ThreadSchemaMutation object of the builder.
func (tsc *ThreadSchemaCreate) Mutation() *ThreadSchemaMutation {
	return tsc.mutation
}

// Save creates the ThreadSchema in the database.
func (tsc *ThreadSchemaCreate) Save(ctx context.Context) (*ThreadSchema, error) {
	return withHooks(ctx, tsc.sqlSave, tsc.mutation, tsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tsc *ThreadSchemaCreate) SaveX(ctx context.Context) *ThreadSchema {
	v, err := tsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tsc *ThreadSchemaCreate) Exec(ctx context.Context) error {
	_, err := tsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsc *ThreadSchemaCreate) ExecX(ctx context.Context) {
	if err := tsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsc *ThreadSchemaCreate) check() error {
	return nil
}

func (tsc *ThreadSchemaCreate) sqlSave(ctx context.Context) (*ThreadSchema, error) {
	if err := tsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tsc.mutation.id = &_node.ID
	tsc.mutation.done = true
	return _node, nil
}

func (tsc *ThreadSchemaCreate) createSpec() (*ThreadSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &ThreadSchema{config: tsc.config}
		_spec = sqlgraph.NewCreateSpec(threadschema.Table, sqlgraph.NewFieldSpec(threadschema.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// ThreadSchemaCreateBulk is the builder for creating many ThreadSchema entities in bulk.
type ThreadSchemaCreateBulk struct {
	config
	builders []*ThreadSchemaCreate
}

// Save creates the ThreadSchema entities in the database.
func (tscb *ThreadSchemaCreateBulk) Save(ctx context.Context) ([]*ThreadSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tscb.builders))
	nodes := make([]*ThreadSchema, len(tscb.builders))
	mutators := make([]Mutator, len(tscb.builders))
	for i := range tscb.builders {
		func(i int, root context.Context) {
			builder := tscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThreadSchemaMutation)
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
					_, err = mutators[i+1].Mutate(root, tscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tscb *ThreadSchemaCreateBulk) SaveX(ctx context.Context) []*ThreadSchema {
	v, err := tscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tscb *ThreadSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := tscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tscb *ThreadSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := tscb.Exec(ctx); err != nil {
		panic(err)
	}
}
