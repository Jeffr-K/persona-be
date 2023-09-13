package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ApplicationSchema holds the schema definition for the ApplicationSchema entity.
type ApplicationSchema struct {
	ent.Schema
}

// Fields of the ApplicationSchema.
func (ApplicationSchema) Fields() []ent.Field {
	return nil
}

// Edges of the ApplicationSchema.
func (ApplicationSchema) Edges() []ent.Edge {
	return nil
}

func (ApplicationSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Application"},
	}
}
