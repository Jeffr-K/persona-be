package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// PositionSchema holds the schema definition for the PositionSchema entity.
type PositionSchema struct {
	ent.Schema
}

// Fields of the PositionSchema.
func (PositionSchema) Fields() []ent.Field {
	return nil
}

// Edges of the PositionSchema.
func (PositionSchema) Edges() []ent.Edge {
	return nil
}

func (PositionSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Position"},
	}
}
