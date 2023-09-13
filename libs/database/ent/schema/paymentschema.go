package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// PaymentSchema holds the schema definition for the PaymentSchema entity.
type PaymentSchema struct {
	ent.Schema
}

// Fields of the PaymentSchema.
func (PaymentSchema) Fields() []ent.Field {
	return nil
}

// Edges of the PaymentSchema.
func (PaymentSchema) Edges() []ent.Edge {
	return nil
}

func (PaymentSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Payment"},
	}
}
