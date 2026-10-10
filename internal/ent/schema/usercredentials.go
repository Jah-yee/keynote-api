package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type UserCredential struct {
	ent.Schema
}

func (UserCredential) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id",uuid.UUID{}).Default(uuid.New),
		field.String("password_hash").NotEmpty().Sensitive(),
	}
}

func (UserCredential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users",User.Type).Ref("credentials").Unique().Required(),
	}
}