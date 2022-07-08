package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Users"},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("TelegramId").Unique(),
		field.String("ActiveStable").Optional(),
		field.String("AllStables").Optional(),
		field.Time("RegTime"),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
