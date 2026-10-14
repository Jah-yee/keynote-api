package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Workspace struct {
	ent.Schema
}

func (Workspace) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("tenant_id",uuid.UUID{}),
		field.String("name").NotEmpty(),
		field.String("slug").NotEmpty(),
		field.UUID("status_id", uuid.UUID{}),
		field.UUID("created_by", uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Workspace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("status",StateDefinitions.Type).Ref("workspace").Field("status_id").Unique().Required(),
		edge.From("tenant",Tenant.Type).Ref("workspace").Field("tenant_id").Unique().Required(),
		edge.From("createdby",User.Type).Ref("creator").Field("created_by").Unique().Required(),
		edge.To("projects",Project.Type),
	}
}

