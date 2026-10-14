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
	"github.com/google/uuid"
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

type SessionAndTokens struct {
	Token string
	RefreshToken string 
	ExpireAt time.Time
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

	out, err := CreateSessionAndTokens(ctx,user.ID,s.db,s.signer)

	output.AccessToken = out.Token
	output.RefreshToken = out.RefreshToken
	output.ExpireAT = out.ExpireAt
	output.RefreshTokenExpireAt = out.RefreshTokenExpireAt

	return output, nil
}

func CreateSessionAndTokens(
	ctx context.Context, 
	userID uuid.UUID,
	db *gen.Client, 
	signer *tokens.Signer,
) (SessionAndTokens,error) {

	var output SessionAndTokens

	refresh, refreshHash, refreshttl, err := tokens.NewRefreshToken()
	if err != nil {
		return output, appError.InvalidCredentials()
	}

	session, err := db.UserSession.Create().
		SetUserID(userID).
		SetRefreshTokenHash(refreshHash).
		SetExpiresAt(time.Now().Add(time.Duration(refreshttl))).
		Save(ctx)
		
	if err != nil {
		return output, appError.InvalidCredentials()
	}

	token, exp, err := signer.Sign(userID, session.ID)
	if err != nil {
		return output, fmt.Errorf("Login error : %v", err)
	}

	output.ExpireAt = exp
	output.Token = token
	output.RefreshToken = refresh
	output.RefreshTokenExpireAt = refreshttl

	return output,nil 
}