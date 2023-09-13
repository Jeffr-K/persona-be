package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// PhoneSchema holds the schema definition for the PhoneSchema entity.
type PhoneSchema struct {
	ent.Schema
}

// Fields of the PhoneSchema.
func (PhoneSchema) Fields() []ent.Field {
	return nil
}

// Edges of the PhoneSchema.
func (PhoneSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", ProfileSchema.Type).
			Ref("phone").
			Unique().
			Required(),
	}
}

func (PhoneSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Phone"},
	}
}
