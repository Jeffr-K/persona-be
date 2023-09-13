// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"persona/libs/database/ent/locationschema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LocationSchemaCreate is the builder for creating a LocationSchema entity.
type LocationSchemaCreate struct {
	config
	mutation *LocationSchemaMutation
	hooks    []Hook
}

// Mutation returns the LocationSchemaMutation object of the builder.
func (lsc *LocationSchemaCreate) Mutation() *LocationSchemaMutation {
	return lsc.mutation
}

// Save creates the LocationSchema in the database.
func (lsc *LocationSchemaCreate) Save(ctx context.Context) (*LocationSchema, error) {
	return withHooks(ctx, lsc.sqlSave, lsc.mutation, lsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lsc *LocationSchemaCreate) SaveX(ctx context.Context) *LocationSchema {
	v, err := lsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lsc *LocationSchemaCreate) Exec(ctx context.Context) error {
	_, err := lsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lsc *LocationSchemaCreate) ExecX(ctx context.Context) {
	if err := lsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lsc *LocationSchemaCreate) check() error {
	return nil
}

func (lsc *LocationSchemaCreate) sqlSave(ctx context.Context) (*LocationSchema, error) {
	if err := lsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lsc.mutation.id = &_node.ID
	lsc.mutation.done = true
	return _node, nil
}

func (lsc *LocationSchemaCreate) createSpec() (*LocationSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &LocationSchema{config: lsc.config}
		_spec = sqlgraph.NewCreateSpec(locationschema.Table, sqlgraph.NewFieldSpec(locationschema.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// LocationSchemaCreateBulk is the builder for creating many LocationSchema entities in bulk.
type LocationSchemaCreateBulk struct {
	config
	builders []*LocationSchemaCreate
}

// Save creates the LocationSchema entities in the database.
func (lscb *LocationSchemaCreateBulk) Save(ctx context.Context) ([]*LocationSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lscb.builders))
	nodes := make([]*LocationSchema, len(lscb.builders))
	mutators := make([]Mutator, len(lscb.builders))
	for i := range lscb.builders {
		func(i int, root context.Context) {
			builder := lscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LocationSchemaMutation)
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
					_, err = mutators[i+1].Mutate(root, lscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lscb *LocationSchemaCreateBulk) SaveX(ctx context.Context) []*LocationSchema {
	v, err := lscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lscb *LocationSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := lscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lscb *LocationSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := lscb.Exec(ctx); err != nil {
		panic(err)
	}
}