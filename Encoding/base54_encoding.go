package Encoding

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

var encoder = map[string] interface{}  {
	"base32": base32.StdEncoding,
	"base64": base64.StdEncoding,
}

func EncodeToBase(data, encodingType string) (string, error) {
	var encodedString string
	enc, exists := encoder[encodingType]
	if !exists {
		return "", fmt.Errorf("invalid encoding type")
	}

	switch enc := enc.(type) {
	case *base32.Encoding:
		encodedString = enc.EncodeToString([]byte(data))
	case *base64.Encoding:
		encodedString = enc.EncodeToString([]byte(data))
	default:
		return "", fmt.Errorf("failed to decode")
	}

	return encodedString, nil
}

func DecodeFromBase(data, encodingType string) (string, error) {
	var decodedStr []byte
	var err error

	enc, exists := encoder[encodingType]
	if !exists {
		return "", fmt.Errorf("invalid encoding type")
	}

	switch enc := enc.(type) {
	case *base32.Encoding:
		decodedStr, err = enc.DecodeString(data)
	case *base64.Encoding:
		decodedStr, err = enc.DecodeString(data)
	default:
		return "", fmt.Errorf("failed to decode")
	}

	if err != nil {
		return "", fmt.Errorf("failed to encode the data: %v", err)
	}
	return string(decodedStr), nil
}