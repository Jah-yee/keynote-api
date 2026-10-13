package auth

import (
	"github.com/danielgtaylor/huma/v2"

	"github.com/caneel7/keynote-api/internal/common/tokens"
	"github.com/caneel7/keynote-api/internal/ent/gen"
	"github.com/caneel7/keynote-api/internal/features/auth/jwks"
	"github.com/caneel7/keynote-api/internal/features/auth/login"
	"github.com/caneel7/keynote-api/internal/features/auth/me"
	"github.com/caneel7/keynote-api/internal/features/auth/register"
)


func Mount(
	public huma.API,
	protected huma.API,
	db *gen.Client,
	signer *tokens.Signer,
) {
	register.RegisterUserRoute(public, register.NewService(db))
	jwks.JwksRoute(public, signer)
	login.LoginRoute(public, login.NewService(db, signer))

	me.GetMeRoute(protected, me.NewService(db))
}
