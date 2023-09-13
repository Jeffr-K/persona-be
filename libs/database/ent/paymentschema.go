// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"persona/libs/database/ent/paymentschema"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PaymentSchema is the model entity for the PaymentSchema schema.
type PaymentSchema struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PaymentSchema) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case paymentschema.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PaymentSchema fields.
func (ps *PaymentSchema) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case paymentschema.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ps.ID = int(value.Int64)
		default:
			ps.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PaymentSchema.
// This includes values selected through modifiers, order, etc.
func (ps *PaymentSchema) Value(name string) (ent.Value, error) {
	return ps.selectValues.Get(name)
}

// Update returns a builder for updating this PaymentSchema.
// Note that you need to call PaymentSchema.Unwrap() before calling this method if this PaymentSchema
// was returned from a transaction, and the transaction was committed or rolled back.
func (ps *PaymentSchema) Update() *PaymentSchemaUpdateOne {
	return NewPaymentSchemaClient(ps.config).UpdateOne(ps)
}

// Unwrap unwraps the PaymentSchema entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ps *PaymentSchema) Unwrap() *PaymentSchema {
	_tx, ok := ps.config.driver.(*txDriver)
	if !ok {
		panic("ent: PaymentSchema is not a transactional entity")
	}
	ps.config.driver = _tx.drv
	return ps
}

// String implements the fmt.Stringer.
func (ps *PaymentSchema) String() string {
	var builder strings.Builder
	builder.WriteString("PaymentSchema(")
	builder.WriteString(fmt.Sprintf("id=%v", ps.ID))
	builder.WriteByte(')')
	return builder.String()
}

// PaymentSchemas is a parsable slice of PaymentSchema.
type PaymentSchemas []*PaymentSchema
