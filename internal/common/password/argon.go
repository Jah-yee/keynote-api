package common

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

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

func VerifyPassword(raw,encoded string) bool {

	parts := strings.Split(encoded,"$")

	if len(parts) != 6 || parts[1] != "argon2id" {
		return false
	}

	var version int
	if _, err := fmt.Sscanf(parts[2],"v=%d",&version); err != nil || version != argon2.Version {
		return false
	}

	var memory, iterations uint32
	var parallelism uint8

	if _,err := fmt.Sscanf(parts[3],"m=%d,t=%d,p=%d",&memory,&iterations,&parallelism); err != nil {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	other := argon2.IDKey([]byte(raw),salt,iterations,memory,parallelism,uint32(len(hash)))
	return subtle.ConstantTimeCompare(hash,other) == 1
}