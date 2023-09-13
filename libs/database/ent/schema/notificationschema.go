package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// NotificationSchema holds the schema definition for the NotificationSchema entity.
type NotificationSchema struct {
	ent.Schema
}

// Fields of the NotificationSchema.
func (NotificationSchema) Fields() []ent.Field {
	return nil
}

// Edges of the NotificationSchema.
func (NotificationSchema) Edges() []ent.Edge {
	return nil
}

func (NotificationSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Notification"},
	}
}
