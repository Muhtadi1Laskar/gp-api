package ciphers

import (
	"strings"
	"unicode"
)

func Rot13Cipher(data string) string {
	var result strings.Builder

	for _, c := range data {
		if unicode.IsLetter(c) {
			var offset rune
			if unicode.IsLower(c) {
				offset = 'a'
			} else {
				offset = 'A'
			}
			roatedString := (c - offset + 13) % 26 + offset
			result.WriteRune(roatedString)
		} else {
			result.WriteRune(c)
		}
	}
	return result.String()
}