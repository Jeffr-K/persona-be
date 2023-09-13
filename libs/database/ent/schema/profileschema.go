package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// ProfileSchema holds the schema definition for the ProfileSchema entity.
type ProfileSchema struct {
	ent.Schema
}

// Fields of the ProfileSchema.
func (ProfileSchema) Fields() []ent.Field {
	return nil
}

// Edges of the ProfileSchema.
func (ProfileSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", UserSchema.Type).Ref("profile").Unique().Required(),
		edge.To("photo", PhotoSchema.Type).Unique(),
		edge.To("phone", PhoneSchema.Type).Unique(),
	}
}

func (ProfileSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Profile"},
	}
}
