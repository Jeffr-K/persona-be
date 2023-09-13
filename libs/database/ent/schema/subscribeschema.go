package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// SubscribeSchema holds the schema definition for the SubscribeSchema entity.
type SubscribeSchema struct {
	ent.Schema
}

// Fields of the SubscribeSchema.
func (SubscribeSchema) Fields() []ent.Field {
	return nil
}

// Edges of the SubscribeSchema.
func (SubscribeSchema) Edges() []ent.Edge {
	return nil
}

func (SubscribeSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Subscribe"},
	}
}
