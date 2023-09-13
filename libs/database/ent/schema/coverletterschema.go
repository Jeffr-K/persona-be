package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// CoverLetterSchema holds the schema definition for the CoverLetterSchema entity.
type CoverLetterSchema struct {
	ent.Schema
}

// Fields of the CoverLetterSchema.
func (CoverLetterSchema) Fields() []ent.Field {
	return nil
}

// Edges of the CoverLetterSchema.
func (CoverLetterSchema) Edges() []ent.Edge {
	return nil
}

func (CoverLetterSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "CoverLetter"},
	}
}
