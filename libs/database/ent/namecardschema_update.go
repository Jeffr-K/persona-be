// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"persona/libs/database/ent/namecardschema"
	"persona/libs/database/ent/predicate"
	"persona/libs/database/ent/userschema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NamecardSchemaUpdate is the builder for updating NamecardSchema entities.
type NamecardSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *NamecardSchemaMutation
}

// Where appends a list predicates to the NamecardSchemaUpdate builder.
func (nsu *NamecardSchemaUpdate) Where(ps ...predicate.NamecardSchema) *NamecardSchemaUpdate {
	nsu.mutation.Where(ps...)
	return nsu
}

// SetUserID sets the "user" edge to the UserSchema entity by ID.
func (nsu *NamecardSchemaUpdate) SetUserID(id int) *NamecardSchemaUpdate {
	nsu.mutation.SetUserID(id)
	return nsu
}

// SetNillableUserID sets the "user" edge to the UserSchema entity by ID if the given value is not nil.
func (nsu *NamecardSchemaUpdate) SetNillableUserID(id *int) *NamecardSchemaUpdate {
	if id != nil {
		nsu = nsu.SetUserID(*id)
	}
	return nsu
}

// SetUser sets the "user" edge to the UserSchema entity.
func (nsu *NamecardSchemaUpdate) SetUser(u *UserSchema) *NamecardSchemaUpdate {
	return nsu.SetUserID(u.ID)
}

// Mutation returns the NamecardSchemaMutation object of the builder.
func (nsu *NamecardSchemaUpdate) Mutation() *NamecardSchemaMutation {
	return nsu.mutation
}

// ClearUser clears the "user" edge to the UserSchema entity.
func (nsu *NamecardSchemaUpdate) ClearUser() *NamecardSchemaUpdate {
	nsu.mutation.ClearUser()
	return nsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nsu *NamecardSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nsu.sqlSave, nsu.mutation, nsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nsu *NamecardSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := nsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nsu *NamecardSchemaUpdate) Exec(ctx context.Context) error {
	_, err := nsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsu *NamecardSchemaUpdate) ExecX(ctx context.Context) {
	if err := nsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nsu *NamecardSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(namecardschema.Table, namecardschema.Columns, sqlgraph.NewFieldSpec(namecardschema.FieldID, field.TypeInt))
	if ps := nsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nsu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   namecardschema.UserTable,
			Columns: []string{namecardschema.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   namecardschema.UserTable,
			Columns: []string{namecardschema.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namecardschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nsu.mutation.done = true
	return n, nil
}

// NamecardSchemaUpdateOne is the builder for updating a single NamecardSchema entity.
type NamecardSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NamecardSchemaMutation
}

// SetUserID sets the "user" edge to the UserSchema entity by ID.
func (nsuo *NamecardSchemaUpdateOne) SetUserID(id int) *NamecardSchemaUpdateOne {
	nsuo.mutation.SetUserID(id)
	return nsuo
}

// SetNillableUserID sets the "user" edge to the UserSchema entity by ID if the given value is not nil.
func (nsuo *NamecardSchemaUpdateOne) SetNillableUserID(id *int) *NamecardSchemaUpdateOne {
	if id != nil {
		nsuo = nsuo.SetUserID(*id)
	}
	return nsuo
}

// SetUser sets the "user" edge to the UserSchema entity.
func (nsuo *NamecardSchemaUpdateOne) SetUser(u *UserSchema) *NamecardSchemaUpdateOne {
	return nsuo.SetUserID(u.ID)
}

// Mutation returns the NamecardSchemaMutation object of the builder.
func (nsuo *NamecardSchemaUpdateOne) Mutation() *NamecardSchemaMutation {
	return nsuo.mutation
}

// ClearUser clears the "user" edge to the UserSchema entity.
func (nsuo *NamecardSchemaUpdateOne) ClearUser() *NamecardSchemaUpdateOne {
	nsuo.mutation.ClearUser()
	return nsuo
}

// Where appends a list predicates to the NamecardSchemaUpdate builder.
func (nsuo *NamecardSchemaUpdateOne) Where(ps ...predicate.NamecardSchema) *NamecardSchemaUpdateOne {
	nsuo.mutation.Where(ps...)
	return nsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nsuo *NamecardSchemaUpdateOne) Select(field string, fields ...string) *NamecardSchemaUpdateOne {
	nsuo.fields = append([]string{field}, fields...)
	return nsuo
}

// Save executes the query and returns the updated NamecardSchema entity.
func (nsuo *NamecardSchemaUpdateOne) Save(ctx context.Context) (*NamecardSchema, error) {
	return withHooks(ctx, nsuo.sqlSave, nsuo.mutation, nsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nsuo *NamecardSchemaUpdateOne) SaveX(ctx context.Context) *NamecardSchema {
	node, err := nsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nsuo *NamecardSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := nsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsuo *NamecardSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := nsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nsuo *NamecardSchemaUpdateOne) sqlSave(ctx context.Context) (_node *NamecardSchema, err error) {
	_spec := sqlgraph.NewUpdateSpec(namecardschema.Table, namecardschema.Columns, sqlgraph.NewFieldSpec(namecardschema.FieldID, field.TypeInt))
	id, ok := nsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NamecardSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, namecardschema.FieldID)
		for _, f := range fields {
			if !namecardschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != namecardschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nsuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   namecardschema.UserTable,
			Columns: []string{namecardschema.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nsuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   namecardschema.UserTable,
			Columns: []string{namecardschema.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NamecardSchema{config: nsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namecardschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nsuo.mutation.done = true
	return _node, nil
}