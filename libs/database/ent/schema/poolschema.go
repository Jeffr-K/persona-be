//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/generated

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// Pool holds the schema definition for the Pool entity.
type PoolSchema struct {
	ent.Schema
}

// Fields of the Pool.
func (PoolSchema) Fields() []ent.Field {
	return nil
}

// Edges of the Pool.
func (PoolSchema) Edges() []ent.Edge {
	return nil
}

func (PoolSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Pool"},
	}
}
