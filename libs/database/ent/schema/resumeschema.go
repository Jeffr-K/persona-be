package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ResumeSchema holds the schema definition for the ResumeSchema entity.
type ResumeSchema struct {
	ent.Schema
}

// Fields of the ResumeSchema.
func (ResumeSchema) Fields() []ent.Field {
	return nil
}

// Edges of the ResumeSchema.
func (ResumeSchema) Edges() []ent.Edge {
	return nil
}

func (ResumeSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Resume"},
	}
}
