package Encoding

import (
	"fmt"
	"strconv"
	"strings"
)

func TextToBinary(data string) string {
    var result strings.Builder

    for _, char := range data {
        result.WriteString(fmt.Sprintf("%08b", char))
    }

    return strings.TrimSpace(result.String())
}

func BinaryToText(data string) (string, error) {
    var result strings.Builder

    if len(data) % 8 != 0 {
        return "", fmt.Errorf("binary string length must be multiple of 8")
    }

    for i := 0; i < len(data); i+=8 {
        byteStr := data[i:i+8]
        charCode, err := strconv.ParseInt(byteStr, 2, 64)
        if err != nil {
            return "", fmt.Errorf("invalid binary chunk: %v", err)
        }
        result.WriteRune(rune(charCode))
    }
    return result.String(), nil
}