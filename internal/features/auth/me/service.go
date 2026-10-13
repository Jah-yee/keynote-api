package me

import (
	"context"
	"fmt"

	"github.com/caneel7/keynote-api/internal/ent/gen"
	"github.com/caneel7/keynote-api/internal/ent/gen/user"
	"github.com/caneel7/keynote-api/internal/errors"
	"github.com/caneel7/keynote-api/internal/types"
	"github.com/google/uuid"
)

type Service struct {
	db *gen.Client
}

func NewService(db *gen.Client) *Service {
	return &Service{db}
}

func (s *Service) GetMe(ctx context.Context, userID uuid.UUID) (types.User, error) {

	u, err := s.db.User.Query().
		Where(user.ID(userID)).
		Only(ctx)
	if err != nil {
		if gen.IsNotFound(err) {
			return types.User{}, errors.New(errors.KindNotFound, "Cannot find user")
		}
		return types.User{}, fmt.Errorf("Error fetching user : %v", err)
	}

	return types.User{
		ID:    u.ID,
		Email: u.Email,
	}, nil
}
