package argon2id

import (
	"crypto/subtle"

	"github.com/zeebo/blake3"
	"golang.org/x/crypto/argon2"
)

const (
	iterations  = 4
	memoryUsage = 64 * (0x1 << 10) // 64MB or 65536KB
	parallelism = 4
	hashLength  = 64
)

func generateSalt(username string) []byte {
	salt := blake3.Sum512([]byte(username))
	return salt[:]
}

func Hash(username, password string) []byte {
	salt := generateSalt(username)
	return argon2.IDKey([]byte(password), salt, iterations, memoryUsage, parallelism, hashLength)
}

func IsMatched(username, password string, hash []byte) bool {
	hashed := Hash(username, password)
	return subtle.ConstantTimeCompare(hashed, hash) == 1
}
