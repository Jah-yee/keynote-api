package auth

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/caneel7/keynote-api/internal/ent/gen"
	"github.com/caneel7/keynote-api/internal/features/auth/register"
)

func Mount(api huma.API, db *gen.Client) {
	register.RegisterUserRoute(api, register.NewService(db))	
}
