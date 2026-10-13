package jwks

import (
	"context"
	"net/http"

	"github.com/caneel7/keynote-api/internal/common/tokens"
	"github.com/danielgtaylor/huma/v2"
)

type Response struct {
	CacheControl string `header:"Cache-Control"`
	Body         tokens.JWKS
}

func JwksRoute(api huma.API, signer *tokens.Signer) {
	huma.Register(
		api,
		huma.Operation{
			OperationID: "jwks",
			Method:      http.MethodGet,
			Path:        "/.well-known/jwks.json",
			Tags:        []string{"Auth"},
			Summary:     "Public signing keys",
		},
		func (ctx context.Context,_ *struct{}) (*Response,error) {
			return &Response{
				CacheControl: "public, max-age=300",
				Body: signer.JWKS(),
			},nil
		},
	)
}
