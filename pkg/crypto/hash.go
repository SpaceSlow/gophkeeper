package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// GenerateHash uses pbkdf-2 via sha256 hash algorithm,
// returns hash with layout: "pbkdf2-sha256$Base64(salt)$passwordIterationsNum$Base64(hash)"
func GenerateHash(password string, keyLen, passwordIterationsNum int) (string, error) {
	salt := make([]byte, keyLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	pbkdf2key := pbkdf2.Key([]byte(password), salt, passwordIterationsNum, keyLen, sha256.New)
	return strings.Join(
		[]string{
			"pbkdf2-sha256",
			base64.StdEncoding.EncodeToString(salt),
			strconv.Itoa(passwordIterationsNum),
			base64.StdEncoding.EncodeToString(pbkdf2key),
		},
		"$",
	), nil
}
