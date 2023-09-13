package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// PhotoSchema holds the schema definition for the PhotoSchema entity.
type PhotoSchema struct {
	ent.Schema
}

// Fields of the PhotoSchema.
func (PhotoSchema) Fields() []ent.Field {
	return nil
}

// Edges of the PhotoSchema.
func (PhotoSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", ProfileSchema.Type).
			Ref("photo").
			Unique().
			Required(),
	}
}

func (PhotoSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "photo"},
	}
}
