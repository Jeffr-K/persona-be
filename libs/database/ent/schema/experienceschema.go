package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ExperienceSchema holds the schema definition for the ExperienceSchema entity.
type ExperienceSchema struct {
	ent.Schema
}

// Fields of the ExperienceSchema.
func (ExperienceSchema) Fields() []ent.Field {
	return nil
}

// Edges of the ExperienceSchema.
func (ExperienceSchema) Edges() []ent.Edge {
	return nil
}

func (ExperienceSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Experience"},
	}
}
