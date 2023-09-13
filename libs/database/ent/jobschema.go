// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"persona/libs/database/ent/jobschema"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// JobSchema is the model entity for the JobSchema schema.
type JobSchema struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JobSchema) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jobschema.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JobSchema fields.
func (js *JobSchema) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jobschema.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			js.ID = int(value.Int64)
		default:
			js.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the JobSchema.
// This includes values selected through modifiers, order, etc.
func (js *JobSchema) Value(name string) (ent.Value, error) {
	return js.selectValues.Get(name)
}

// Update returns a builder for updating this JobSchema.
// Note that you need to call JobSchema.Unwrap() before calling this method if this JobSchema
// was returned from a transaction, and the transaction was committed or rolled back.
func (js *JobSchema) Update() *JobSchemaUpdateOne {
	return NewJobSchemaClient(js.config).UpdateOne(js)
}

// Unwrap unwraps the JobSchema entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (js *JobSchema) Unwrap() *JobSchema {
	_tx, ok := js.config.driver.(*txDriver)
	if !ok {
		panic("ent: JobSchema is not a transactional entity")
	}
	js.config.driver = _tx.drv
	return js
}

// String implements the fmt.Stringer.
func (js *JobSchema) String() string {
	var builder strings.Builder
	builder.WriteString("JobSchema(")
	builder.WriteString(fmt.Sprintf("id=%v", js.ID))
	builder.WriteByte(')')
	return builder.String()
}

// JobSchemas is a parsable slice of JobSchema.
type JobSchemas []*JobSchema
