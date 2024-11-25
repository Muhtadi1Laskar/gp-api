package hashes

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashData(message string) string {
	hash := sha256.New()
	hash.Write([]byte(message))

	hashedBytes := hash.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}
