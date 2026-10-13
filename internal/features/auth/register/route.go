package register

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	httpError "github.com/caneel7/keynote-api/internal/errors"
	types "github.com/caneel7/keynote-api/internal/types"
)

type Registrar interface {
	Register(
		ctx context.Context,
		email string,
		password string,
	) (types.User, error)
}

type RegisterUserRequest struct {
	Body struct {
		Email    string `json:"email" format:"email" maxLength:"254"`
		Password string `json:"password" minLenght:"8" maxLength:"128"`
	}
}

type RegisterUserResponse struct {
	Body struct {
		ID    string `json:"id" format:"uuid" doc:"Registered user's generated ID"`
		Email string `json:"email" format:"email" doc:"Registered user's email address"`
	}
}

func RegisterUserRoute(api huma.API,svc Registrar) {
	huma.Register(
		api,
		huma.Operation{
			OperationID: "register",
			Method: http.MethodPost,
			Path: "/auth/register",
			DefaultStatus: http.StatusCreated,
			Tags: []string{"Auth"},
			Summary: "Register new user",
			Description: "Register new user in system",
		},
		func (ctx context.Context,req *RegisterUserRequest) (*RegisterUserResponse,error) {
			u, err := svc.Register(ctx,req.Body.Email,req.Body.Password)
			if err != nil {
				return nil, httpError.From(err)
			}
			res := &RegisterUserResponse{}
			res.Body.Email = u.Email
			res.Body.ID = u.ID.String()
			return res,nil
		},
	)
}