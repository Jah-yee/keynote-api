package tokens

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var b64 = base64.RawURLEncoding.EncodeToString

type signingKey struct {
	kid  string
	priv *rsa.PrivateKey
}

type Signer struct {
	keys            []signingKey
	issuer          string
	audience        string
	ttl             time.Duration
	refreshTokenTTL time.Duration
}

type Claims struct {
	SessionID string `json:"sid"`
	jwt.RegisteredClaims
}

var refreshTokenTTl time.Duration

func NewSigner(issuer, audience string, refreshTokenttl, ttl time.Duration, privs ...*rsa.PrivateKey) (*Signer, error) {

	if len(privs) == 0 {
		return nil, fmt.Errorf("Invalid Private key : %v", errors.New("tokens : at least one key required"))
	}

	s := &Signer{
		issuer:   issuer,
		audience: audience,
		ttl:      ttl,
	}

	refreshTokenTTl = refreshTokenttl

	for _, p := range privs {
		if p.N.BitLen() < 2048 {
			return nil, fmt.Errorf("Invalid Private key : %v", errors.New("tokens: RSA key must be at least 2048 bits"))
		}
		s.keys = append(s.keys, signingKey{kid: thumbprint(&p.PublicKey), priv: p})
	}
	return s, nil
}

func ParsePrivateKeyPEM(data []byte) (*rsa.PrivateKey, error) {

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("Invalid Private key : %v", errors.New("tokens: no PEM block found"))
	}

	if k, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		rk, ok := k.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("Invalid Private key : %v", errors.New("tokens: key is not RSA"))
		}
		return rk, nil
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func thumbprint(pub *rsa.PublicKey) string {
	canon := fmt.Sprintf(`{"e":"%s","kty":"RSA","n":"%s"}`,
		b64(big.NewInt(int64(pub.E)).Bytes()), b64(pub.N.Bytes()))
	sum := sha256.Sum256([]byte(canon))
	return b64(sum[:])
}

func (s *Signer) Sign(userID, sid uuid.UUID) (string, time.Time, error) {

	now := time.Now()
	exp := now.Add(s.ttl)
	t := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		Claims{
			SessionID: sid.String(),
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    s.issuer,
				Subject:   userID.String(),
				Audience:  jwt.ClaimStrings{s.audience},
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(exp),
				ID:        uuid.NewString(),
			},
		},
	)

	t.Header["kid"] = s.keys[0].kid
	signed, err := t.SignedString(s.keys[0].priv)
	return signed, exp, err
}

type JWK struct {
	Kty string `json:"kty"`
	Use string `json:"user"`
	Alg string `json:"alg"`
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

func (s *Signer) JWKS() JWKS {

	out := JWKS{Keys: make([]JWK, 0, len(s.keys))}

	for _, k := range s.keys {
		out.Keys = append(out.Keys, JWK{
			Kty: "RSA", Use: "sig", Alg: "RS256", Kid: k.kid,
			N: b64(k.priv.N.Bytes()),
			E: b64(big.NewInt(int64(k.priv.E)).Bytes()),
		})
	}
	return out
}

func NewRefreshToken() (plain, hash string, ttl time.Duration, err error) {
	raw := make([]byte, 32)
	if _, err = rand.Read(raw); err != nil {
		return "", "",refreshTokenTTl, fmt.Errorf("read random: %w", err)
	}
	plain = base64.RawURLEncoding.EncodeToString(raw)
	sum := sha256.Sum256([]byte(plain))
	return plain, hex.EncodeToString(sum[:]),refreshTokenTTl, nil
}
