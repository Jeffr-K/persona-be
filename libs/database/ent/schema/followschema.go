package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// FollowSchema holds the schema definition for the FollowSchema entity.
type FollowSchema struct {
	ent.Schema
}

// Fields of the FollowSchema.
func (FollowSchema) Fields() []ent.Field {
	return nil
}

// Edges of the FollowSchema.
func (FollowSchema) Edges() []ent.Edge {
	return nil
}

func (FollowSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Follow"},
	}
}
