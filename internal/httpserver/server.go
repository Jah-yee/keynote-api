package httpserver

import (
	"net"
	"net/http"

	"github.com/caneel7/keynote-api/config"
	"github.com/caneel7/keynote-api/internal/ent/gen"
)

func NewServer (c config.HTTP, db *gen.Client) *http.Server {
	return &http.Server{
		Addr: net.JoinHostPort("",c.Addr),
		Handler: routes(),
		ReadTimeout: c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	}
}