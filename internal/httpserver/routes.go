package httpserver

import (
	"net/http"

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

func routes () http.Handler {

	router := chi.NewRouter()

	healthHandler(newApi(router))

	return router;
}