// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"persona/libs/database/ent/jobschema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobSchemaCreate is the builder for creating a JobSchema entity.
type JobSchemaCreate struct {
	config
	mutation *JobSchemaMutation
	hooks    []Hook
}

// Mutation returns the JobSchemaMutation object of the builder.
func (jsc *JobSchemaCreate) Mutation() *JobSchemaMutation {
	return jsc.mutation
}

// Save creates the JobSchema in the database.
func (jsc *JobSchemaCreate) Save(ctx context.Context) (*JobSchema, error) {
	return withHooks(ctx, jsc.sqlSave, jsc.mutation, jsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jsc *JobSchemaCreate) SaveX(ctx context.Context) *JobSchema {
	v, err := jsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jsc *JobSchemaCreate) Exec(ctx context.Context) error {
	_, err := jsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jsc *JobSchemaCreate) ExecX(ctx context.Context) {
	if err := jsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jsc *JobSchemaCreate) check() error {
	return nil
}

func (jsc *JobSchemaCreate) sqlSave(ctx context.Context) (*JobSchema, error) {
	if err := jsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jsc.mutation.id = &_node.ID
	jsc.mutation.done = true
	return _node, nil
}

func (jsc *JobSchemaCreate) createSpec() (*JobSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &JobSchema{config: jsc.config}
		_spec = sqlgraph.NewCreateSpec(jobschema.Table, sqlgraph.NewFieldSpec(jobschema.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// JobSchemaCreateBulk is the builder for creating many JobSchema entities in bulk.
type JobSchemaCreateBulk struct {
	config
	builders []*JobSchemaCreate
}

// Save creates the JobSchema entities in the database.
func (jscb *JobSchemaCreateBulk) Save(ctx context.Context) ([]*JobSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(jscb.builders))
	nodes := make([]*JobSchema, len(jscb.builders))
	mutators := make([]Mutator, len(jscb.builders))
	for i := range jscb.builders {
		func(i int, root context.Context) {
			builder := jscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobSchemaMutation)
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
					_, err = mutators[i+1].Mutate(root, jscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, jscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jscb *JobSchemaCreateBulk) SaveX(ctx context.Context) []*JobSchema {
	v, err := jscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jscb *JobSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := jscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jscb *JobSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := jscb.Exec(ctx); err != nil {
		panic(err)
	}
}
