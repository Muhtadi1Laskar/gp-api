package ciphers

import (
	"strings"
	"unicode"
)

func EncryptCaeser(data string, key int) string {
	var result strings.Builder
	key = key % 26
	if key < 0 {
			key += 26
		}

	for _, char := range data {
		if unicode.IsLetter(char) {
			offset := 'A'
			if unicode.IsLower(char) {
				offset = 'a'
			}
			// Perform the shift and wrap around the alphabet.
			newChar := (char-offset+rune(key)+26)%26 + offset
			result.WriteRune(newChar)
		} else {
			// Leave non-alphabetic characters as they are.
			result.WriteRune(char)
		}
	}
	return result.String()
}