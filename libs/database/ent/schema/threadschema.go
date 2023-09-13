package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ThreadSchema holds the schema definition for the ThreadSchema entity.
type ThreadSchema struct {
	ent.Schema
}

// Fields of the ThreadSchema.
func (ThreadSchema) Fields() []ent.Field {
	return nil
}

// Edges of the ThreadSchema.
func (ThreadSchema) Edges() []ent.Edge {
	return nil
}

func (ThreadSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Thread"},
	}
}
