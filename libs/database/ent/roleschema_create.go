// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"persona/libs/database/ent/roleschema"
	"persona/libs/database/ent/userschema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleSchemaCreate is the builder for creating a RoleSchema entity.
type RoleSchemaCreate struct {
	config
	mutation *RoleSchemaMutation
	hooks    []Hook
}

// SetUserID sets the "user" edge to the UserSchema entity by ID.
func (rsc *RoleSchemaCreate) SetUserID(id int) *RoleSchemaCreate {
	rsc.mutation.SetUserID(id)
	return rsc
}

// SetNillableUserID sets the "user" edge to the UserSchema entity by ID if the given value is not nil.
func (rsc *RoleSchemaCreate) SetNillableUserID(id *int) *RoleSchemaCreate {
	if id != nil {
		rsc = rsc.SetUserID(*id)
	}
	return rsc
}

// SetUser sets the "user" edge to the UserSchema entity.
func (rsc *RoleSchemaCreate) SetUser(u *UserSchema) *RoleSchemaCreate {
	return rsc.SetUserID(u.ID)
}

// Mutation returns the RoleSchemaMutation object of the builder.
func (rsc *RoleSchemaCreate) Mutation() *RoleSchemaMutation {
	return rsc.mutation
}

// Save creates the RoleSchema in the database.
func (rsc *RoleSchemaCreate) Save(ctx context.Context) (*RoleSchema, error) {
	return withHooks(ctx, rsc.sqlSave, rsc.mutation, rsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rsc *RoleSchemaCreate) SaveX(ctx context.Context) *RoleSchema {
	v, err := rsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rsc *RoleSchemaCreate) Exec(ctx context.Context) error {
	_, err := rsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsc *RoleSchemaCreate) ExecX(ctx context.Context) {
	if err := rsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rsc *RoleSchemaCreate) check() error {
	return nil
}

func (rsc *RoleSchemaCreate) sqlSave(ctx context.Context) (*RoleSchema, error) {
	if err := rsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rsc.mutation.id = &_node.ID
	rsc.mutation.done = true
	return _node, nil
}

func (rsc *RoleSchemaCreate) createSpec() (*RoleSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleSchema{config: rsc.config}
		_spec = sqlgraph.NewCreateSpec(roleschema.Table, sqlgraph.NewFieldSpec(roleschema.FieldID, field.TypeInt))
	)
	if nodes := rsc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roleschema.UserTable,
			Columns: []string{roleschema.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_schema_roles = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoleSchemaCreateBulk is the builder for creating many RoleSchema entities in bulk.
type RoleSchemaCreateBulk struct {
	config
	builders []*RoleSchemaCreate
}

// Save creates the RoleSchema entities in the database.
func (rscb *RoleSchemaCreateBulk) Save(ctx context.Context) ([]*RoleSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rscb.builders))
	nodes := make([]*RoleSchema, len(rscb.builders))
	mutators := make([]Mutator, len(rscb.builders))
	for i := range rscb.builders {
		func(i int, root context.Context) {
			builder := rscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleSchemaMutation)
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
					_, err = mutators[i+1].Mutate(root, rscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rscb *RoleSchemaCreateBulk) SaveX(ctx context.Context) []*RoleSchema {
	v, err := rscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rscb *RoleSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := rscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rscb *RoleSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := rscb.Exec(ctx); err != nil {
		panic(err)
	}
}