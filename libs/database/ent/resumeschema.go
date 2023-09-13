// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"persona/libs/database/ent/resumeschema"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ResumeSchema is the model entity for the ResumeSchema schema.
type ResumeSchema struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResumeSchema) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resumeschema.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResumeSchema fields.
func (rs *ResumeSchema) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resumeschema.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rs.ID = int(value.Int64)
		default:
			rs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ResumeSchema.
// This includes values selected through modifiers, order, etc.
func (rs *ResumeSchema) Value(name string) (ent.Value, error) {
	return rs.selectValues.Get(name)
}

// Update returns a builder for updating this ResumeSchema.
// Note that you need to call ResumeSchema.Unwrap() before calling this method if this ResumeSchema
// was returned from a transaction, and the transaction was committed or rolled back.
func (rs *ResumeSchema) Update() *ResumeSchemaUpdateOne {
	return NewResumeSchemaClient(rs.config).UpdateOne(rs)
}

// Unwrap unwraps the ResumeSchema entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rs *ResumeSchema) Unwrap() *ResumeSchema {
	_tx, ok := rs.config.driver.(*txDriver)
	if !ok {
		panic("ent: ResumeSchema is not a transactional entity")
	}
	rs.config.driver = _tx.drv
	return rs
}

// String implements the fmt.Stringer.
func (rs *ResumeSchema) String() string {
	var builder strings.Builder
	builder.WriteString("ResumeSchema(")
	builder.WriteString(fmt.Sprintf("id=%v", rs.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ResumeSchemas is a parsable slice of ResumeSchema.
type ResumeSchemas []*ResumeSchema
