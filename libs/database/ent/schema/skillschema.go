package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// SkillSchema holds the schema definition for the SkillSchema entity.
type SkillSchema struct {
	ent.Schema
}

// Fields of the SkillSchema.
func (SkillSchema) Fields() []ent.Field {
	return nil
}

// Edges of the SkillSchema.
func (SkillSchema) Edges() []ent.Edge {
	return nil
}

func (SkillSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Skill"},
	}
}
