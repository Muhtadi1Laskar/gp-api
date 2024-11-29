package ciphers

import (
	"strings"
	"unicode"
)

func AtbashCipher(data string) string {
	var result strings.Builder

	for _, c := range data {
		if unicode.IsLetter(c) {
			var offset rune
			if unicode.IsLower(c) {
				offset = 'a'
			} else {
				offset = 'A'
			}
			reversedRune := rune(offset + ('Z' - c))
			result.WriteRune(reversedRune)
		} else {
			result.WriteRune(c)
		}
	}
	return result.String()
}