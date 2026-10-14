package tenant

import (
	"github.com/caneel7/keynote-api/internal/ent/gen"
	"github.com/caneel7/keynote-api/internal/features/tenant/create"
	"github.com/danielgtaylor/huma/v2"
)

func Mount(
	protected huma.API,
	db *gen.Client,
) {
	/*
	* @path : "/tenant"
	*/
	create.CreateTenantRoute(protected, create.NewService(db))

}
