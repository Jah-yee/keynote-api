package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type UserSession struct{ ent.Schema }

func (UserSession) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id",uuid.UUID{}),
		field.String("refresh_token_hash").NotEmpty().Unique().Sensitive(),
		field.String("user_agent").Optional().MaxLen(512),
		field.Time("expires_at"),
		field.Time("last_used_at").Optional().Nillable(),
		field.Time("revoked_at").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (UserSession) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("sessions").Field("user_id").Unique().Required(),
	}
}

func (UserSession) Indexes() []ent.Index {
	return []ent.Index{index.Edges("user")} // "list this user's sessions"
}