package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// BadgeSchema holds the schema definition for the BadgeSchema entity.
type BadgeSchema struct {
	ent.Schema
}

// Fields of the BadgeSchema.
func (BadgeSchema) Fields() []ent.Field {
	return nil
}

// Edges of the BadgeSchema.
func (BadgeSchema) Edges() []ent.Edge {
	return nil
}

func (BadgeSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Badge"},
	}
}
