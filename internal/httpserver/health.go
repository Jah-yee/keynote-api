package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

type Pinger interface {
	PingContext(ctx context.Context) error
}


type ServerStatusResponse struct {
	Body struct {
		Status string `json:"status" example:"Ok" doc:"Server Status"`
		Timestamp string `json:"timestamp" format:"datet-time" doc:"Server time in utc"`
	}
}

func newServerStatusResponseBuilder (status string) *ServerStatusResponse {
	response := &ServerStatusResponse{}
	response.Body.Status = status
	response.Body.Timestamp = time.Now().UTC().Format(time.RFC3339)
	return response
}


func healthHandler (api huma.API) {
	huma.Register(
		api,
		huma.Operation{
			OperationID: "health",
			Method: http.MethodGet,
			Path: "/health",
			Tags: []string{"System"},
			Summary: "Server Status Check",
			Description: "Returns Server Status in UTC Time",
		},
		func (ctx context.Context, _ *struct{}) (*ServerStatusResponse,error) {
			return newServerStatusResponseBuilder("Ok"),nil
		},
	)
}