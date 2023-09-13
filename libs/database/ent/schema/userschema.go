package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type UserSchema struct {
	ent.Schema
}

func (UserSchema) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New).Comment("UUID"),
		field.String("username").Comment("사용자 이름"),
		field.String("password").Comment("사용자 비밀번호"),
		field.String("email").Unique().Comment("사용자 이메일"),
		field.Time("createdAt").Default(time.Now().UTC).Comment("사용자 등록 일시"),
		field.Time("updatedAt").Default(time.Now().UTC).Comment("사용자 수정 일시"),
	}
}

func (UserSchema) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", RoleSchema.Type),
		edge.To("profile", ProfileSchema.Type).Unique(),
		edge.To("follow", FollowSchema.Type).Unique(),
		edge.To("referrer", ReferrerSchema.Type).Unique(),
		edge.To("personalization", PersonalizationSchema.Type).Unique(),
		edge.To("namecard", NamecardSchema.Type).Unique(),
	}
}

func (UserSchema) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "User"},
	}
}
