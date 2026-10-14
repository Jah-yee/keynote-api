package slug

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("[^a-z0-9]+")

func Randomsuffix() string {
	b := make([]byte, 3)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func Slugify(s string) string {
	// 1. Lowercase the string
	s = strings.ToLower(s)
	// 2. Replace non-alphanumeric characters with hyphens
	s = re.ReplaceAllString(s, "-")
	// 3. Trim leading and trailing hyphens
	return strings.Trim(s, "-")
}