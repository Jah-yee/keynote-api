package login

import (
	"context"
	"fmt"
	"time"

	argon "github.com/caneel7/keynote-api/internal/common/password"
	"github.com/caneel7/keynote-api/internal/common/tokens"
	"github.com/caneel7/keynote-api/internal/ent/gen"
	"github.com/caneel7/keynote-api/internal/ent/gen/user"
	"github.com/caneel7/keynote-api/internal/ent/gen/usercredential"
	appError "github.com/caneel7/keynote-api/internal/errors"
)

type Service struct {
	db     *gen.Client
	signer *tokens.Signer
}

func NewService(db *gen.Client, signer *tokens.Signer) *Service {
	return &Service{db: db, signer: signer}
}

type LoginOutput struct {
	AccessToken          string
	RefreshToken         string
	ExpireAT             time.Time
	RefreshTokenExpireAt time.Duration
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (LoginOutput, error) {

	var output LoginOutput

	u, err := s.db.UserCredential.Query().
		Where(usercredential.HasUsersWith(user.Email(email))).
		WithUsers().
		Only(ctx)


	if err != nil {
		if gen.IsNotFound(err) {
			return output, appError.InvalidCredentials()
		}
		return output, fmt.Errorf("Login error : %v", err)
	}

	user := u.Edges.Users

	ok := argon.VerifyPassword(password, u.PasswordHash)
	if !ok {
		return output, appError.InvalidCredentials()
	}

	refresh, refreshHash, refreshttl, err := tokens.NewRefreshToken()
	if err != nil {
		return output, appError.InvalidCredentials()
	}

	session, err := s.db.UserSession.Create().
		SetUserID(user.ID).
		SetRefreshTokenHash(refreshHash).
		SetExpiresAt(time.Now().Add(time.Duration(refreshttl))).
		Save(ctx)
		
	if err != nil {
		return output, appError.InvalidCredentials()
	}

	token, exp, err := s.signer.Sign(user.ID, session.ID)
	if err != nil {
		return output, fmt.Errorf("Login error : %v", err)
	}

	output.AccessToken = token
	output.RefreshToken = refresh
	output.ExpireAT = exp
	output.RefreshTokenExpireAt = refreshttl

	return output, nil
}
