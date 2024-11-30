package ciphers

import (
	"unicode"
)

func formatKey(key string) string {
	formatted := make([]rune, len(key))

	for _, char := range key {
		if unicode.IsLetter(char) {
			formatted = append(formatted, unicode.ToUpper(char))
		}
	}
	return string(formatted)
}

func VigenereCipher(data, key string, encrypt bool) string {
	key = formatKey(key)
	keyShifts := make([]int, len(key))
	for i, r := range key {
		keyShifts[i] = int(r - 'A')
	}

	result := make([]rune, len(data))

	for i, j := 0, 0; i < len(data); i++ {
		char := rune(data[i])
		if !unicode.IsLetter(char) {
			result[i] = char
			continue
		}

		shift := keyShifts[j%len(keyShifts)]
		if !encrypt {
			shift = -shift
		}

		if unicode.IsUpper(char) {
			result[i] = 'A' + (char-'A'+rune(shift)+26)%26
		} else {
			result[i] = 'a' + (char-'a'+rune(shift)+26)%26
		}
		j++
	}

	return string(result)
}
