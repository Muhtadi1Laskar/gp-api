package ciphers

import (
	"unicode"
)

func VigenereCipher(data, key string, encrypt bool) string {
	keyShifts := make([]int, len(key))
	for i, r := range key {
		if unicode.IsUpper(r) {
			keyShifts[i] = int(r - 'A')
		} else if unicode.IsLower(r) {
			keyShifts[i] = int(r - 'a')
		} else {
			keyShifts[i] = 0
		}
	}

	result := make([]rune, len(data))

	for i, r := range data {
		if !unicode.IsLetter(r) {
			result[r] = r
			continue
		}

		shift := keyShifts[i % len(keyShifts)]
		if !encrypt {
			shift = -shift
		}

		if unicode.IsUpper(r) {
			result[i] = 'A' + (r - 'A' + rune(shift) + 26) % 26
		} else if unicode.IsLower(r) {
			result[i] = 'a' + (r - 'a' + rune(shift) + 26) % 26
		}
	}
	return string(result)
}