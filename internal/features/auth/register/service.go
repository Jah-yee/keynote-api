package register

import (
	"context"
	"fmt"
	"net/mail"
	"strings"

	argon "github.com/caneel7/keynote-api/internal/common/password"
	"github.com/caneel7/keynote-api/internal/ent/gen"
	appError "github.com/caneel7/keynote-api/internal/errors"
	types "github.com/caneel7/keynote-api/internal/types"
)


type Service struct {
	db *gen.Client
}

func NewService(db *gen.Client) *Service {
	return &Service{db}
}

func (s *Service) Register(
	ctx context.Context,
	email string,
	password string,
) (types.User, error) {

	email = strings.ToLower(strings.TrimSpace(email))

	if _, err := mail.ParseAddress(email); err != nil {
		return types.User{}, appError.Invalid(map[string]string{
			email: "Invalid Email",
		}) 
	}

	hash, err := argon.HashPassword(password);
	if err != nil {
		return types.User{}, err
	}

	tx, err := s.db.Tx(ctx)
	if err != nil {
		return types.User{},nil
	}
	defer tx.Rollback()

	user, err := tx.User.Create().SetEmail(email).Save(ctx)
	if err != nil {
		if gen.IsConstraintError(err) {
			return types.User{}, appError.New(
				appError.KindConflict,
				"Email Already Exists",
			)
		}
		return types.User{},fmt.Errorf("Error creating user : %v",err)
	}

	if _, err := tx.UserCredential.Create().SetUsers(user).SetPasswordHash(hash).Save(ctx); err != nil {
		return types.User{},fmt.Errorf("Error saving user credentials : %v",err)
	}

	if err := tx.Commit(); err != nil {
		return types.User{}, fmt.Errorf("Error commiting to db : %v", err)
	}

	return types.User{
		ID: user.ID,
		Email: user.Email,
	},nil
}
