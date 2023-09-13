package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// PortfolioSchema holds the schema definition for the PortfolioSchema entity.
type PortfolioSchema struct {
	ent.Schema
}

// Fields of the PortfolioSchema.
func (PortfolioSchema) Fields() []ent.Field {
	return nil
}

// Edges of the PortfolioSchema.
func (PortfolioSchema) Edges() []ent.Edge {
	return nil
}

func (PortfolioSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Portfolio"},
	}
}
