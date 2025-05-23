package hashes

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func GetHashFunc() map[string]func() (hash.Hash, error) {
	return map[string]func() (hash.Hash, error){
		// MD Family
		"md4": func() (hash.Hash, error) { return md4.New(), nil },
		"md5": func() (hash.Hash, error) { return md5.New(), nil },

		// SHA-1 Family
		"sha1": func() (hash.Hash, error) { return sha1.New(), nil },

		// SHA-2 Family
		"sha256":     func() (hash.Hash, error) { return sha256.New(), nil },
		"sha224":     func() (hash.Hash, error) { return sha256.New224(), nil },
		"sha384":     func() (hash.Hash, error) { return sha512.New384(), nil },
		"sha512_224": func() (hash.Hash, error) { return sha512.New512_224(), nil },
		"sha512":     func() (hash.Hash, error) { return sha512.New(), nil },
		"sha512_256": func() (hash.Hash, error) { return sha512.New512_256(), nil },

		// SHA-3 Family
		"sha3_224": func() (hash.Hash, error) { return sha3.New224(), nil },
		"sha3_256": func() (hash.Hash, error) { return sha3.New256(), nil },
		"sha3_384": func() (hash.Hash, error) { return sha3.New384(), nil },
		"sha3_512": func() (hash.Hash, error) { return sha3.New512(), nil },

		// RIPEMD
		"ripemd160": func() (hash.Hash, error) { return ripemd160.New(), nil },

		// BLAKE2 Family
		"blake2s_256": func() (hash.Hash, error) { return blake2s.New256(nil) },
		"blake2b_256": func() (hash.Hash, error) { return blake2b.New256(nil) },
		"blake2b_384": func() (hash.Hash, error) { return blake2b.New384(nil) },
		"blake2b_512": func() (hash.Hash, error) { return blake2b.New512(nil) },
	}
}

func HashData(message, hashName string) (string, error) {
	hashList := GetHashFunc()
	hashFunc, exists := hashList[strings.ToLower(hashName)]
	if !exists {
		return "", fmt.Errorf("unsupported hash algorithm: %s", hashName)
	}

	hasher, err := hashFunc()
	if err != nil {
		return "", fmt.Errorf("failed to create hash instance: %v", err)
	}

	hasher.Write([]byte(message))

	hashedBytes := hasher.Sum(nil)

	return hex.EncodeToString(hashedBytes), nil
}

func VerifyHash(originalHash, message, hashName string) (bool, error) {
	newHash, err := HashData(message, hashName)
	if err != nil {
		return false, err
	}

	return newHash == originalHash, nil
}
