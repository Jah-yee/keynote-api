package httpserver

import (
	"net"
	"net/http"

	"github.com/caneel7/keynote-api/config"
)

func NewServer (c config.HTTP) *http.Server {
	return &http.Server{
		Addr: net.JoinHostPort("",c.Addr),
		Handler: routes(),
		ReadTimeout: c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	}
}