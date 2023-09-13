package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// LogoSchema holds the schema definition for the LogoSchema entity.
type LogoSchema struct {
	ent.Schema
}

// Fields of the LogoSchema.
func (LogoSchema) Fields() []ent.Field {
	return nil
}

// Edges of the LogoSchema.
func (LogoSchema) Edges() []ent.Edge {
	return nil
}

func (LogoSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Logo"},
	}
}
