package httpserver

import (
	"net/http"

	auth "github.com/caneel7/keynote-api/internal/features/auth"
	"github.com/caneel7/keynote-api/internal/ent/gen"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func newApi (router chi.Router) huma.API {
	return humachi.New(
		router,
		huma.DefaultConfig("keynote-api","0.1.0"),
	)
}

func routes (db *gen.Client) http.Handler {

	router := chi.NewRouter()
	api := newApi(router)

	healthHandler(api)
	auth.Mount(api,db)

	return router;
}