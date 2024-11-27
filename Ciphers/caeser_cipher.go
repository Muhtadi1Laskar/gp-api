package ciphers

func EncryptCaeser(data string, key int) string {
	var result []rune

	key = key % 26
	if key < 0 {
		key += 26
	}

	for _, char := range data {
		if char >= 'A' && char <= 'Z' {
			result = append(result, rune('A'+(char-'A'+rune(key))%26))
		} else if char >= 'a' && char <= 'z' {
			result = append(result, rune('a'+(char-'a'+rune(key))%26))
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
