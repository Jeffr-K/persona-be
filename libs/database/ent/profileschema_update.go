// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"persona/libs/database/ent/phoneschema"
	"persona/libs/database/ent/photoschema"
	"persona/libs/database/ent/predicate"
	"persona/libs/database/ent/profileschema"
	"persona/libs/database/ent/userschema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileSchemaUpdate is the builder for updating ProfileSchema entities.
type ProfileSchemaUpdate struct {
	config
	hooks    []Hook
	mutation *ProfileSchemaMutation
}

// Where appends a list predicates to the ProfileSchemaUpdate builder.
func (psu *ProfileSchemaUpdate) Where(ps ...predicate.ProfileSchema) *ProfileSchemaUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetOwnerID sets the "owner" edge to the UserSchema entity by ID.
func (psu *ProfileSchemaUpdate) SetOwnerID(id int) *ProfileSchemaUpdate {
	psu.mutation.SetOwnerID(id)
	return psu
}

// SetOwner sets the "owner" edge to the UserSchema entity.
func (psu *ProfileSchemaUpdate) SetOwner(u *UserSchema) *ProfileSchemaUpdate {
	return psu.SetOwnerID(u.ID)
}

// SetPhotoID sets the "photo" edge to the PhotoSchema entity by ID.
func (psu *ProfileSchemaUpdate) SetPhotoID(id int) *ProfileSchemaUpdate {
	psu.mutation.SetPhotoID(id)
	return psu
}

// SetNillablePhotoID sets the "photo" edge to the PhotoSchema entity by ID if the given value is not nil.
func (psu *ProfileSchemaUpdate) SetNillablePhotoID(id *int) *ProfileSchemaUpdate {
	if id != nil {
		psu = psu.SetPhotoID(*id)
	}
	return psu
}

// SetPhoto sets the "photo" edge to the PhotoSchema entity.
func (psu *ProfileSchemaUpdate) SetPhoto(p *PhotoSchema) *ProfileSchemaUpdate {
	return psu.SetPhotoID(p.ID)
}

// SetPhoneID sets the "phone" edge to the PhoneSchema entity by ID.
func (psu *ProfileSchemaUpdate) SetPhoneID(id int) *ProfileSchemaUpdate {
	psu.mutation.SetPhoneID(id)
	return psu
}

// SetNillablePhoneID sets the "phone" edge to the PhoneSchema entity by ID if the given value is not nil.
func (psu *ProfileSchemaUpdate) SetNillablePhoneID(id *int) *ProfileSchemaUpdate {
	if id != nil {
		psu = psu.SetPhoneID(*id)
	}
	return psu
}

// SetPhone sets the "phone" edge to the PhoneSchema entity.
func (psu *ProfileSchemaUpdate) SetPhone(p *PhoneSchema) *ProfileSchemaUpdate {
	return psu.SetPhoneID(p.ID)
}

// Mutation returns the ProfileSchemaMutation object of the builder.
func (psu *ProfileSchemaUpdate) Mutation() *ProfileSchemaMutation {
	return psu.mutation
}

// ClearOwner clears the "owner" edge to the UserSchema entity.
func (psu *ProfileSchemaUpdate) ClearOwner() *ProfileSchemaUpdate {
	psu.mutation.ClearOwner()
	return psu
}

// ClearPhoto clears the "photo" edge to the PhotoSchema entity.
func (psu *ProfileSchemaUpdate) ClearPhoto() *ProfileSchemaUpdate {
	psu.mutation.ClearPhoto()
	return psu
}

// ClearPhone clears the "phone" edge to the PhoneSchema entity.
func (psu *ProfileSchemaUpdate) ClearPhone() *ProfileSchemaUpdate {
	psu.mutation.ClearPhone()
	return psu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *ProfileSchemaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, psu.sqlSave, psu.mutation, psu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psu *ProfileSchemaUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *ProfileSchemaUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *ProfileSchemaUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psu *ProfileSchemaUpdate) check() error {
	if _, ok := psu.mutation.OwnerID(); psu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProfileSchema.owner"`)
	}
	return nil
}

func (psu *ProfileSchemaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := psu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(profileschema.Table, profileschema.Columns, sqlgraph.NewFieldSpec(profileschema.FieldID, field.TypeInt))
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if psu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profileschema.OwnerTable,
			Columns: []string{profileschema.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profileschema.OwnerTable,
			Columns: []string{profileschema.OwnerColumn},
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
	if psu.mutation.PhotoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profileschema.PhotoTable,
			Columns: []string{profileschema.PhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(photoschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.PhotoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profileschema.PhotoTable,
			Columns: []string{profileschema.PhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(photoschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psu.mutation.PhoneCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profileschema.PhoneTable,
			Columns: []string{profileschema.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(phoneschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.PhoneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profileschema.PhoneTable,
			Columns: []string{profileschema.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(phoneschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profileschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	psu.mutation.done = true
	return n, nil
}

// ProfileSchemaUpdateOne is the builder for updating a single ProfileSchema entity.
type ProfileSchemaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProfileSchemaMutation
}

// SetOwnerID sets the "owner" edge to the UserSchema entity by ID.
func (psuo *ProfileSchemaUpdateOne) SetOwnerID(id int) *ProfileSchemaUpdateOne {
	psuo.mutation.SetOwnerID(id)
	return psuo
}

// SetOwner sets the "owner" edge to the UserSchema entity.
func (psuo *ProfileSchemaUpdateOne) SetOwner(u *UserSchema) *ProfileSchemaUpdateOne {
	return psuo.SetOwnerID(u.ID)
}

// SetPhotoID sets the "photo" edge to the PhotoSchema entity by ID.
func (psuo *ProfileSchemaUpdateOne) SetPhotoID(id int) *ProfileSchemaUpdateOne {
	psuo.mutation.SetPhotoID(id)
	return psuo
}

// SetNillablePhotoID sets the "photo" edge to the PhotoSchema entity by ID if the given value is not nil.
func (psuo *ProfileSchemaUpdateOne) SetNillablePhotoID(id *int) *ProfileSchemaUpdateOne {
	if id != nil {
		psuo = psuo.SetPhotoID(*id)
	}
	return psuo
}

// SetPhoto sets the "photo" edge to the PhotoSchema entity.
func (psuo *ProfileSchemaUpdateOne) SetPhoto(p *PhotoSchema) *ProfileSchemaUpdateOne {
	return psuo.SetPhotoID(p.ID)
}

// SetPhoneID sets the "phone" edge to the PhoneSchema entity by ID.
func (psuo *ProfileSchemaUpdateOne) SetPhoneID(id int) *ProfileSchemaUpdateOne {
	psuo.mutation.SetPhoneID(id)
	return psuo
}

// SetNillablePhoneID sets the "phone" edge to the PhoneSchema entity by ID if the given value is not nil.
func (psuo *ProfileSchemaUpdateOne) SetNillablePhoneID(id *int) *ProfileSchemaUpdateOne {
	if id != nil {
		psuo = psuo.SetPhoneID(*id)
	}
	return psuo
}

// SetPhone sets the "phone" edge to the PhoneSchema entity.
func (psuo *ProfileSchemaUpdateOne) SetPhone(p *PhoneSchema) *ProfileSchemaUpdateOne {
	return psuo.SetPhoneID(p.ID)
}

// Mutation returns the ProfileSchemaMutation object of the builder.
func (psuo *ProfileSchemaUpdateOne) Mutation() *ProfileSchemaMutation {
	return psuo.mutation
}

// ClearOwner clears the "owner" edge to the UserSchema entity.
func (psuo *ProfileSchemaUpdateOne) ClearOwner() *ProfileSchemaUpdateOne {
	psuo.mutation.ClearOwner()
	return psuo
}

// ClearPhoto clears the "photo" edge to the PhotoSchema entity.
func (psuo *ProfileSchemaUpdateOne) ClearPhoto() *ProfileSchemaUpdateOne {
	psuo.mutation.ClearPhoto()
	return psuo
}

// ClearPhone clears the "phone" edge to the PhoneSchema entity.
func (psuo *ProfileSchemaUpdateOne) ClearPhone() *ProfileSchemaUpdateOne {
	psuo.mutation.ClearPhone()
	return psuo
}

// Where appends a list predicates to the ProfileSchemaUpdate builder.
func (psuo *ProfileSchemaUpdateOne) Where(ps ...predicate.ProfileSchema) *ProfileSchemaUpdateOne {
	psuo.mutation.Where(ps...)
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *ProfileSchemaUpdateOne) Select(field string, fields ...string) *ProfileSchemaUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated ProfileSchema entity.
func (psuo *ProfileSchemaUpdateOne) Save(ctx context.Context) (*ProfileSchema, error) {
	return withHooks(ctx, psuo.sqlSave, psuo.mutation, psuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *ProfileSchemaUpdateOne) SaveX(ctx context.Context) *ProfileSchema {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *ProfileSchemaUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *ProfileSchemaUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psuo *ProfileSchemaUpdateOne) check() error {
	if _, ok := psuo.mutation.OwnerID(); psuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProfileSchema.owner"`)
	}
	return nil
}

func (psuo *ProfileSchemaUpdateOne) sqlSave(ctx context.Context) (_node *ProfileSchema, err error) {
	if err := psuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(profileschema.Table, profileschema.Columns, sqlgraph.NewFieldSpec(profileschema.FieldID, field.TypeInt))
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProfileSchema.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profileschema.FieldID)
		for _, f := range fields {
			if !profileschema.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profileschema.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if psuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profileschema.OwnerTable,
			Columns: []string{profileschema.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   profileschema.OwnerTable,
			Columns: []string{profileschema.OwnerColumn},
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
	if psuo.mutation.PhotoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profileschema.PhotoTable,
			Columns: []string{profileschema.PhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(photoschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.PhotoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profileschema.PhotoTable,
			Columns: []string{profileschema.PhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(photoschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psuo.mutation.PhoneCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profileschema.PhoneTable,
			Columns: []string{profileschema.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(phoneschema.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.PhoneIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   profileschema.PhoneTable,
			Columns: []string{profileschema.PhoneColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(phoneschema.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProfileSchema{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profileschema.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	psuo.mutation.done = true
	return _node, nil
}