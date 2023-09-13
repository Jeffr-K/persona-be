package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// LikeSchema holds the schema definition for the LikeSchema entity.
type LikeSchema struct {
	ent.Schema
}

// Fields of the LikeSchema.
func (LikeSchema) Fields() []ent.Field {
	return nil
}

// Edges of the LikeSchema.
func (LikeSchema) Edges() []ent.Edge {
	return nil
}

func (LikeSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Like"},
	}
}
