package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AssignmentSchema holds the schema definition for the AssignmentSchema entity.
type AssignmentSchema struct {
	ent.Schema
}

// Fields of the AssignmentSchema.
func (AssignmentSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type"),
	}
}

// Edges of the AssignmentSchema.
func (AssignmentSchema) Edges() []ent.Edge {
	return nil
}

func (AssignmentSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Assignment"},
	}
}
