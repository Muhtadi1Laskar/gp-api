package ciphers

import (
	"encoding/hex"
	"fmt"
)

func EncryptXOR(plainText, key string) string {
	output := make([]byte, len(plainText))

	for i := 0; i < len(plainText); i++ {
		output[i] = plainText[i] ^ key[i % len(key)]
	}

	return hex.EncodeToString(output)
}

func DecryptXOR(cipherText, key string) (string, error) {
	hexBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("failed to convert hex string to hex bytes: %v", err)
	}

	output := make([]byte, len(hexBytes))

	for i := 0; i < len(hexBytes); i++ {
		output[i] = hexBytes[i] ^ key[i % len(key)]
	}

	return string(output), nil
}
