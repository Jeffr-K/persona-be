package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// PersonalizationSchema holds the schema definition for the PersonalizationSchema entity.
type PersonalizationSchema struct {
	ent.Schema
}

// Fields of the PersonalizationSchema.
func (PersonalizationSchema) Fields() []ent.Field {
	return nil
}

// Edges of the PersonalizationSchema.
func (PersonalizationSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", UserSchema.Type).
			Ref("personalization").
			Unique(),
	}
}

func (PersonalizationSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Personalization"},
	}
}
