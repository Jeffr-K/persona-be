// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"persona/libs/database/ent/badgeschema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BadgeSchemaCreate is the builder for creating a BadgeSchema entity.
type BadgeSchemaCreate struct {
	config
	mutation *BadgeSchemaMutation
	hooks    []Hook
}

// Mutation returns the BadgeSchemaMutation object of the builder.
func (bsc *BadgeSchemaCreate) Mutation() *BadgeSchemaMutation {
	return bsc.mutation
}

// Save creates the BadgeSchema in the database.
func (bsc *BadgeSchemaCreate) Save(ctx context.Context) (*BadgeSchema, error) {
	return withHooks(ctx, bsc.sqlSave, bsc.mutation, bsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bsc *BadgeSchemaCreate) SaveX(ctx context.Context) *BadgeSchema {
	v, err := bsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bsc *BadgeSchemaCreate) Exec(ctx context.Context) error {
	_, err := bsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bsc *BadgeSchemaCreate) ExecX(ctx context.Context) {
	if err := bsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bsc *BadgeSchemaCreate) check() error {
	return nil
}

func (bsc *BadgeSchemaCreate) sqlSave(ctx context.Context) (*BadgeSchema, error) {
	if err := bsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bsc.mutation.id = &_node.ID
	bsc.mutation.done = true
	return _node, nil
}

func (bsc *BadgeSchemaCreate) createSpec() (*BadgeSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &BadgeSchema{config: bsc.config}
		_spec = sqlgraph.NewCreateSpec(badgeschema.Table, sqlgraph.NewFieldSpec(badgeschema.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// BadgeSchemaCreateBulk is the builder for creating many BadgeSchema entities in bulk.
type BadgeSchemaCreateBulk struct {
	config
	builders []*BadgeSchemaCreate
}

// Save creates the BadgeSchema entities in the database.
func (bscb *BadgeSchemaCreateBulk) Save(ctx context.Context) ([]*BadgeSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bscb.builders))
	nodes := make([]*BadgeSchema, len(bscb.builders))
	mutators := make([]Mutator, len(bscb.builders))
	for i := range bscb.builders {
		func(i int, root context.Context) {
			builder := bscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BadgeSchemaMutation)
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
					_, err = mutators[i+1].Mutate(root, bscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bscb *BadgeSchemaCreateBulk) SaveX(ctx context.Context) []*BadgeSchema {
	v, err := bscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bscb *BadgeSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := bscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bscb *BadgeSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := bscb.Exec(ctx); err != nil {
		panic(err)
	}
}