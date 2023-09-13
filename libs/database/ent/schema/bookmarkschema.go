package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// BookmarkSchema holds the schema definition for the BookmarkSchema entity.
type BookmarkSchema struct {
	ent.Schema
}

// Fields of the BookmarkSchema.
func (BookmarkSchema) Fields() []ent.Field {
	return nil
}

// Edges of the BookmarkSchema.
func (BookmarkSchema) Edges() []ent.Edge {
	return nil
}

func (BookmarkSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Bookmark"},
	}
}
