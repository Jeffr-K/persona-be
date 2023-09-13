// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"persona/libs/database/ent/bookmarkschema"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BookmarkSchema is the model entity for the BookmarkSchema schema.
type BookmarkSchema struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BookmarkSchema) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookmarkschema.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BookmarkSchema fields.
func (bs *BookmarkSchema) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookmarkschema.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bs.ID = int(value.Int64)
		default:
			bs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BookmarkSchema.
// This includes values selected through modifiers, order, etc.
func (bs *BookmarkSchema) Value(name string) (ent.Value, error) {
	return bs.selectValues.Get(name)
}

// Update returns a builder for updating this BookmarkSchema.
// Note that you need to call BookmarkSchema.Unwrap() before calling this method if this BookmarkSchema
// was returned from a transaction, and the transaction was committed or rolled back.
func (bs *BookmarkSchema) Update() *BookmarkSchemaUpdateOne {
	return NewBookmarkSchemaClient(bs.config).UpdateOne(bs)
}

// Unwrap unwraps the BookmarkSchema entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bs *BookmarkSchema) Unwrap() *BookmarkSchema {
	_tx, ok := bs.config.driver.(*txDriver)
	if !ok {
		panic("ent: BookmarkSchema is not a transactional entity")
	}
	bs.config.driver = _tx.drv
	return bs
}

// String implements the fmt.Stringer.
func (bs *BookmarkSchema) String() string {
	var builder strings.Builder
	builder.WriteString("BookmarkSchema(")
	builder.WriteString(fmt.Sprintf("id=%v", bs.ID))
	builder.WriteByte(')')
	return builder.String()
}

// BookmarkSchemas is a parsable slice of BookmarkSchema.
type BookmarkSchemas []*BookmarkSchema