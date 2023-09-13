package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// CompanySchema holds the schema definition for the CompanySchema entity.
type CompanySchema struct {
	ent.Schema
}

// Fields of the CompanySchema.
func (CompanySchema) Fields() []ent.Field {
	return nil
}

// Edges of the CompanySchema.
func (CompanySchema) Edges() []ent.Edge {
	return nil
}

func (CompanySchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Company"},
	}
}
