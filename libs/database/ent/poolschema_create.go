// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"persona/libs/database/ent/poolschema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PoolSchemaCreate is the builder for creating a PoolSchema entity.
type PoolSchemaCreate struct {
	config
	mutation *PoolSchemaMutation
	hooks    []Hook
}

// Mutation returns the PoolSchemaMutation object of the builder.
func (psc *PoolSchemaCreate) Mutation() *PoolSchemaMutation {
	return psc.mutation
}

// Save creates the PoolSchema in the database.
func (psc *PoolSchemaCreate) Save(ctx context.Context) (*PoolSchema, error) {
	return withHooks(ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PoolSchemaCreate) SaveX(ctx context.Context) *PoolSchema {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *PoolSchemaCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *PoolSchemaCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PoolSchemaCreate) check() error {
	return nil
}

func (psc *PoolSchemaCreate) sqlSave(ctx context.Context) (*PoolSchema, error) {
	if err := psc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	psc.mutation.id = &_node.ID
	psc.mutation.done = true
	return _node, nil
}

func (psc *PoolSchemaCreate) createSpec() (*PoolSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &PoolSchema{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(poolschema.Table, sqlgraph.NewFieldSpec(poolschema.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// PoolSchemaCreateBulk is the builder for creating many PoolSchema entities in bulk.
type PoolSchemaCreateBulk struct {
	config
	builders []*PoolSchemaCreate
}

// Save creates the PoolSchema entities in the database.
func (pscb *PoolSchemaCreateBulk) Save(ctx context.Context) ([]*PoolSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PoolSchema, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PoolSchemaMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *PoolSchemaCreateBulk) SaveX(ctx context.Context) []*PoolSchema {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *PoolSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *PoolSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}
