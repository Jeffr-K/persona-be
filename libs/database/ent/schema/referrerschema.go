package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// ReferrerSchema holds the schema definition for the ReferrerSchema entity.
type ReferrerSchema struct {
	ent.Schema
}

// Fields of the ReferrerSchema.
func (ReferrerSchema) Fields() []ent.Field {
	return nil
}

// Edges of the ReferrerSchema.
func (ReferrerSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", UserSchema.Type).
			Ref("referrer").
			Unique(),
	}
}

func (ReferrerSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Referrer"},
	}
}
