package ciphers

import (
	"strings"
)

var MorseCodes = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.", 'G': "--.", 'H': "....",
	'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---", 'P': ".--.",
	'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..", '0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
	'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.", '.': ".-.-.-", ',': "--..--",
	'?': "..--..", '\'': ".----.", '!': "-.-.--", '/': "-..-.", '(': "-.--.", ')': "-.--.-", '&': ".-...",
	':': "---...", ';': "-.-.-.", '=': "-...-", '+': ".-.-.", '-': "-....-", '_': "..--.-", '"': ".-..-.",
	'$': "...-..-", '@': ".--.-.", ' ': "/",
}

var morseToChar map[string]string

func init() {
	morseToChar := make(map[string]string)
	for char, morse := range MorseCodes {
		morseToChar[morse] = string(char)
	}
}

func getMorseCode(key rune) string {
	if value := MorseCodes[key]; value != "" {
		return value
	}
	return "[?]"
}

func getMorseLetter(str string) rune {
	for key, value := range MorseCodes {
		if value == str {
			return key
		}
	}
	return '?'
}

func MorseCodeEncrypt(data string) string {
	var result strings.Builder
	capitalizedWord := strings.ToUpper(data)

	for _, c := range capitalizedWord {
		if code := getMorseCode(c); code != "[?]" {
			result.WriteString(code + " ")
		} else {
			result.WriteString("[?] ")
		}
	}
	return strings.TrimSpace(result.String())
}

func MorseCodeDecrypt(cipher string) string {
	var result strings.Builder
	formattedStr := strings.Split(cipher, " ")

	for _, c := range formattedStr {
		if letter := getMorseLetter(c); letter != '?' {
			result.WriteRune(letter)
		} else {
			result.WriteRune('?')
		}
	}
	return result.String()
}

