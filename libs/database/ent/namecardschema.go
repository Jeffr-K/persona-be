// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"persona/libs/database/ent/namecardschema"
	"persona/libs/database/ent/userschema"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// NamecardSchema is the model entity for the NamecardSchema schema.
type NamecardSchema struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NamecardSchemaQuery when eager-loading is set.
	Edges                NamecardSchemaEdges `json:"edges"`
	user_schema_namecard *int
	selectValues         sql.SelectValues
}

// NamecardSchemaEdges holds the relations/edges for other nodes in the graph.
type NamecardSchemaEdges struct {
	// User holds the value of the user edge.
	User *UserSchema `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NamecardSchemaEdges) UserOrErr() (*UserSchema, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: userschema.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NamecardSchema) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case namecardschema.FieldID:
			values[i] = new(sql.NullInt64)
		case namecardschema.ForeignKeys[0]: // user_schema_namecard
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NamecardSchema fields.
func (ns *NamecardSchema) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case namecardschema.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ns.ID = int(value.Int64)
		case namecardschema.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_schema_namecard", value)
			} else if value.Valid {
				ns.user_schema_namecard = new(int)
				*ns.user_schema_namecard = int(value.Int64)
			}
		default:
			ns.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NamecardSchema.
// This includes values selected through modifiers, order, etc.
func (ns *NamecardSchema) Value(name string) (ent.Value, error) {
	return ns.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the NamecardSchema entity.
func (ns *NamecardSchema) QueryUser() *UserSchemaQuery {
	return NewNamecardSchemaClient(ns.config).QueryUser(ns)
}

// Update returns a builder for updating this NamecardSchema.
// Note that you need to call NamecardSchema.Unwrap() before calling this method if this NamecardSchema
// was returned from a transaction, and the transaction was committed or rolled back.
func (ns *NamecardSchema) Update() *NamecardSchemaUpdateOne {
	return NewNamecardSchemaClient(ns.config).UpdateOne(ns)
}

// Unwrap unwraps the NamecardSchema entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ns *NamecardSchema) Unwrap() *NamecardSchema {
	_tx, ok := ns.config.driver.(*txDriver)
	if !ok {
		panic("ent: NamecardSchema is not a transactional entity")
	}
	ns.config.driver = _tx.drv
	return ns
}

// String implements the fmt.Stringer.
func (ns *NamecardSchema) String() string {
	var builder strings.Builder
	builder.WriteString("NamecardSchema(")
	builder.WriteString(fmt.Sprintf("id=%v", ns.ID))
	builder.WriteByte(')')
	return builder.String()
}

// NamecardSchemas is a parsable slice of NamecardSchema.
type NamecardSchemas []*NamecardSchema