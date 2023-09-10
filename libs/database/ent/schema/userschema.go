package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type UserSchema struct {
	ent.Schema
}

func (UserSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("password"),
		field.String("email").Unique(),
	}
}

func (UserSchema) Edges() []ent.Edge {
	return nil
}

func (UserSchema) Config() ent.Config {
	return ent.Config{
		Table: "User",
	}
}
