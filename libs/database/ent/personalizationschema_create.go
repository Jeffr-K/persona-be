// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"persona/libs/database/ent/personalizationschema"
	"persona/libs/database/ent/userschema"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PersonalizationSchemaCreate is the builder for creating a PersonalizationSchema entity.
type PersonalizationSchemaCreate struct {
	config
	mutation *PersonalizationSchemaMutation
	hooks    []Hook
}

// SetUserID sets the "user" edge to the UserSchema entity by ID.
func (psc *PersonalizationSchemaCreate) SetUserID(id int) *PersonalizationSchemaCreate {
	psc.mutation.SetUserID(id)
	return psc
}

// SetNillableUserID sets the "user" edge to the UserSchema entity by ID if the given value is not nil.
func (psc *PersonalizationSchemaCreate) SetNillableUserID(id *int) *PersonalizationSchemaCreate {
	if id != nil {
		psc = psc.SetUserID(*id)
	}
	return psc
}

// SetUser sets the "user" edge to the UserSchema entity.
func (psc *PersonalizationSchemaCreate) SetUser(u *UserSchema) *PersonalizationSchemaCreate {
	return psc.SetUserID(u.ID)
}

// Mutation returns the PersonalizationSchemaMutation object of the builder.
func (psc *PersonalizationSchemaCreate) Mutation() *PersonalizationSchemaMutation {
	return psc.mutation
}

// Save creates the PersonalizationSchema in the database.
func (psc *PersonalizationSchemaCreate) Save(ctx context.Context) (*PersonalizationSchema, error) {
	return withHooks(ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PersonalizationSchemaCreate) SaveX(ctx context.Context) *PersonalizationSchema {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *PersonalizationSchemaCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *PersonalizationSchemaCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PersonalizationSchemaCreate) check() error {
	return nil
}

func (psc *PersonalizationSchemaCreate) sqlSave(ctx context.Context) (*PersonalizationSchema, error) {
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

func (psc *PersonalizationSchemaCreate) createSpec() (*PersonalizationSchema, *sqlgraph.CreateSpec) {
	var (
		_node = &PersonalizationSchema{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(personalizationschema.Table, sqlgraph.NewFieldSpec(personalizationschema.FieldID, field.TypeInt))
	)
	if nodes := psc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   personalizationschema.UserTable,
			Columns: []string{personalizationschema.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_schema_personalization = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PersonalizationSchemaCreateBulk is the builder for creating many PersonalizationSchema entities in bulk.
type PersonalizationSchemaCreateBulk struct {
	config
	builders []*PersonalizationSchemaCreate
}

// Save creates the PersonalizationSchema entities in the database.
func (pscb *PersonalizationSchemaCreateBulk) Save(ctx context.Context) ([]*PersonalizationSchema, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PersonalizationSchema, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonalizationSchemaMutation)
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
func (pscb *PersonalizationSchemaCreateBulk) SaveX(ctx context.Context) []*PersonalizationSchema {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *PersonalizationSchemaCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *PersonalizationSchemaCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}