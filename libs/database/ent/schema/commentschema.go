package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// CommentSchema holds the schema definition for the CommentSchema entity.
type CommentSchema struct {
	ent.Schema
}

// Fields of the CommentSchema.
func (CommentSchema) Fields() []ent.Field {
	return nil
}

// Edges of the CommentSchema.
func (CommentSchema) Edges() []ent.Edge {
	return nil
}

func (CommentSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Comment"},
	}
}
