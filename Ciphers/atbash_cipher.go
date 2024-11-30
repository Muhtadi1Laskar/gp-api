package ciphers

import (
	"strings"
	"unicode"
)

func AtbashCipher(data string) string {
	var result strings.Builder

	for _, char := range data {
		if unicode.IsLetter(char) {
			var cipher rune
			if unicode.IsUpper(char) {
				cipher = 'Z' - char + 'A'
			} else {
				cipher = 'z' - char + 'a'
			}
			result.WriteRune(cipher)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}