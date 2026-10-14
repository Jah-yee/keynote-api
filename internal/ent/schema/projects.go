package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Project struct {
	ent.Schema
}

func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id",uuid.UUID{}).Default(uuid.New),
		field.UUID("tenant_id",uuid.UUID{}),
		field.UUID("workspace_id",uuid.UUID{}),
		field.String("name").NotEmpty(),
		field.UUID("status_id",uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant",Tenant.Type).Ref("projects").Field("tenant_id").Unique().Required(),
		edge.From("workspace",Workspace.Type).Ref("projects").Field("workspace_id").Unique().Required(),
		edge.From("state",StateDefinitions.Type).Ref("projects").Field("status_id").Unique().Required(),
	}
}