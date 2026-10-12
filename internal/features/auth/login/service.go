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

type loginOutput struct {
	AccessToken          string
	RefreshToken         string
	ExpireAT             time.Time
	RefreshTokenExpireAt time.Duration
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (loginOutput, error) {

	u, err := s.db.UserCredential.Query().
		Where(usercredential.HasUsersWith(user.Email(email))).
		WithUsers().
		Only(ctx)


	if err != nil {
		if gen.IsNotFound(err) {
			return loginOutput{}, appError.InvalidCredentials()
		}
		return loginOutput{}, fmt.Errorf("Login error : %v", err)
	}

	ok := argon.VerifyPassword(password, u.PasswordHash)
	if !ok {
		return loginOutput{}, appError.InvalidCredentials()
	}

	refresh, refreshHash, refreshttl, err := tokens.NewRefreshToken()
	if err != nil {
		return loginOutput{}, appError.InvalidCredentials()
	}

	session, err := s.db.UserSession.Create().
		SetUserID(u.ID).
		SetRefreshTokenHash(refreshHash).
		SetExpiresAt(time.Now().Add(time.Duration(refreshttl))).
		Save(ctx)

	fmt.Print(err)
		
	if err != nil {
		return loginOutput{}, appError.InvalidCredentials()
	}

	token, exp, err := s.signer.Sign(u.ID, session.ID)
	if err != nil {
		return loginOutput{}, fmt.Errorf("Login error : %v", err)
	}

	return loginOutput{
		AccessToken:  token,
		RefreshToken: refresh,
		ExpireAT:     exp,
		RefreshTokenExpireAt: refreshttl,
	}, nil
}
