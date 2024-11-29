package ciphers

import (
	"errors"
	"strings"
	"unicode"
)

var codes = map[rune]string{'A': "aaaaa", 'B': "aaaab", 'C': "aaaba", 'D': "aaabb", 'E': "aabaa",
	'F': "aabab", 'G': "aabba", 'H': "aabbb", 'I': "abaaa", 'J': "abaab",
	'K': "ababa", 'L': "ababb", 'M': "abbaa", 'N': "abbab", 'O': "abbba",
	'P': "abbbb", 'Q': "baaaa", 'R': "baaab", 'S': "baaba", 'T': "baabb",
	'U': "babaa", 'V': "babab", 'W': "babba", 'X': "babbb", 'Y': "bbaaa", 'Z': "bbaab",
}

func getCode(key rune) string {
	if value, ok := codes[key]; ok {
		return value
	}
	return ""
}

func getLetter(str string) rune {
	for key, value := range codes {
		if value == str {
			return key
		}
	}
	return '#'
}

func BaconCipherEncrypt(data string) string {
	var result strings.Builder
	capitalizedWord := strings.ToUpper(data)

	for _, c := range capitalizedWord {
		if unicode.IsLetter(c) {
			letter := getCode(c)
			result.WriteString(letter + " ")
		}
	}
	return strings.TrimSpace(result.String())
}

func BaconCipherDecrypt(data string) (string, error) {
	var result strings.Builder
	formattedStr := strings.ReplaceAll(data, " ", "")

	if len(formattedStr) % 5 != 0 {
		return "", errors.New("invalid input length")
	}

	for i := 0; i < len(formattedStr); i+= 5 {
		substr := formattedStr[i : i+5]
		letter := string(getLetter(substr))

		result.WriteString(letter)
	}
	return result.String(), nil
}
