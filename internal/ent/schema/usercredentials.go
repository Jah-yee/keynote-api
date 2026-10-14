package schema

import (
	"time"

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
		field.UUID("user_id",uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (UserCredential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users",User.Type).Field("user_id").Ref("credentials").Unique().Required(),
	}
}