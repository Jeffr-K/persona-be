package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// LeetcodeSchema holds the schema definition for the LeetcodeSchema entity.
type LeetcodeSchema struct {
	ent.Schema
}

// Fields of the LeetcodeSchema.
func (LeetcodeSchema) Fields() []ent.Field {
	return nil
}

// Edges of the LeetcodeSchema.
func (LeetcodeSchema) Edges() []ent.Edge {
	return nil
}

func (LeetcodeSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Leetcode"},
	}
}
