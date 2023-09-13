package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// LocationSchema holds the schema definition for the LocationSchema entity.
type LocationSchema struct {
	ent.Schema
}

// Fields of the LocationSchema.
func (LocationSchema) Fields() []ent.Field {
	return nil
}

// Edges of the LocationSchema.
func (LocationSchema) Edges() []ent.Edge {
	return nil
}

func (LocationSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Location"},
	}
}
