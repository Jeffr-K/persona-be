package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// RoleSchema holds the schema definition for the RoleSchema entity.
type RoleSchema struct {
	ent.Schema
}

// Fields of the RoleSchema.
func (RoleSchema) Fields() []ent.Field {
	return nil
}

// Edges of the RoleSchema.
func (RoleSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", UserSchema.Type).
			Ref("roles").
			Unique(),
	}
}

func (RoleSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Role"},
	}
}
