package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ImageSchema holds the schema definition for the ImageSchema entity.
type ImageSchema struct {
	ent.Schema
}

// Fields of the ImageSchema.
func (ImageSchema) Fields() []ent.Field {
	return nil
}

// Edges of the ImageSchema.
func (ImageSchema) Edges() []ent.Edge {
	return nil
}

func (ImageSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Image"},
	}
}
