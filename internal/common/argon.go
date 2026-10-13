package common

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	argonMemoryKiB = 19 * 1024
	argonTime      = 2
	argonThreads   = 1
	argonKeyLen    = 32
	argonSaltLen   = 16
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, argonSaltLen)

	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("Error generating salt : %v", err)
	}

	key := argon2.IDKey(
		[]byte(password),
		salt,
		argonTime,
		argonMemoryKiB,
		argonThreads,
		argonKeyLen,
	)
	b64 := base64.RawStdEncoding

	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		argonMemoryKiB,
		argonTime,
		argonThreads,
		b64.EncodeToString(salt),
		b64.EncodeToString(key),
	), nil

}
