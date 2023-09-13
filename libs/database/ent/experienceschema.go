// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"persona/libs/database/ent/experienceschema"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ExperienceSchema is the model entity for the ExperienceSchema schema.
type ExperienceSchema struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExperienceSchema) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case experienceschema.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExperienceSchema fields.
func (es *ExperienceSchema) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case experienceschema.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			es.ID = int(value.Int64)
		default:
			es.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExperienceSchema.
// This includes values selected through modifiers, order, etc.
func (es *ExperienceSchema) Value(name string) (ent.Value, error) {
	return es.selectValues.Get(name)
}

// Update returns a builder for updating this ExperienceSchema.
// Note that you need to call ExperienceSchema.Unwrap() before calling this method if this ExperienceSchema
// was returned from a transaction, and the transaction was committed or rolled back.
func (es *ExperienceSchema) Update() *ExperienceSchemaUpdateOne {
	return NewExperienceSchemaClient(es.config).UpdateOne(es)
}

// Unwrap unwraps the ExperienceSchema entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (es *ExperienceSchema) Unwrap() *ExperienceSchema {
	_tx, ok := es.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExperienceSchema is not a transactional entity")
	}
	es.config.driver = _tx.drv
	return es
}

// String implements the fmt.Stringer.
func (es *ExperienceSchema) String() string {
	var builder strings.Builder
	builder.WriteString("ExperienceSchema(")
	builder.WriteString(fmt.Sprintf("id=%v", es.ID))
	builder.WriteByte(')')
	return builder.String()
}

// ExperienceSchemas is a parsable slice of ExperienceSchema.
type ExperienceSchemas []*ExperienceSchema
