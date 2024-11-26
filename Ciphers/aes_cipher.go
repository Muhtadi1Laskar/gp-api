package ciphers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func AESEncrypt(plainText, secretKey string) (string, error) {
	aseFunc, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("unable to create the instance of aes cipher: %s", err)
	}

	gcm, err := cipher.NewGCM(aseFunc)
	if err != nil {
		return "", fmt.Errorf("unable to create the gcm: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", fmt.Errorf("unable to generate a random number: %v", err)
	}

	cipherByte := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return hex.EncodeToString(cipherByte), nil
}