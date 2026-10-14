package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)


type User struct {
	ent.Schema
}

func (User) Fields () []ent.Field {
	return []ent.Field {
		field.UUID("id",uuid.UUID{}).Default(uuid.New),
		field.String("email").NotEmpty().Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (User) Edges () []ent.Edge {
	return []ent.Edge {
		edge.To("credentials",UserCredential.Type).Unique(),
		edge.To("sessions",UserSession.Type),
	}
}