package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// NamecardSchema holds the schema definition for the NamecardSchema entity.
type NamecardSchema struct {
	ent.Schema
}

// Fields of the NamecardSchema.
func (NamecardSchema) Fields() []ent.Field {
	return nil
}

// Edges of the NamecardSchema.
func (NamecardSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", UserSchema.Type).
			Ref("namecard").
			Unique(),
	}
}

func (NamecardSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Namecard"},
	}
}
