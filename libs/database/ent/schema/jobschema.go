package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// JobSchema holds the schema definition for the JobSchema entity.
type JobSchema struct {
	ent.Schema
}

// Fields of the JobSchema.
func (JobSchema) Fields() []ent.Field {
	return nil
}

// Edges of the JobSchema.
func (JobSchema) Edges() []ent.Edge {
	return nil
}

func (JobSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Job"},
	}
}
