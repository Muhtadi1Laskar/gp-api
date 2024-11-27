package ciphers

import (
	"encoding/hex"
	"fmt"
)

func EncryptXOR(plainText, key string) string {
	var output string

	for i := 0; i < len(plainText); i++ {
		output += string(plainText[i] ^ key[i%len(key)])
	}
	return hex.EncodeToString([]byte(output))
}

func DecryptXOR(cipherText, key string) (string, error) {
	hexBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("failed to convert hex string to hex bytes: %v", err)
	}
	var plainText string = string(hexBytes)
	var output string

	for i := 0; i < len(plainText); i++ {
		output += string(plainText[i] ^ key[i%len(key)]) 
	}
	return output, nil
}
