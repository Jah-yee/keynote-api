package auth

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/caneel7/keynote-api/internal/common/tokens"
	"github.com/caneel7/keynote-api/internal/ent/gen"
	"github.com/caneel7/keynote-api/internal/features/auth/jwks"
	"github.com/caneel7/keynote-api/internal/features/auth/login"
	"github.com/caneel7/keynote-api/internal/features/auth/register"
)

func Mount(api huma.API, db *gen.Client,signer *tokens.Signer) {
	register.RegisterUserRoute(api, register.NewService(db))
	jwks.JwksRoute(api,signer)
	login.LoginRoute(api,login.NewService(db,signer))
}
