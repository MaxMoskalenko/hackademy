package cipher

import "regexp"

type Vigenere struct {
	key string
}

func (cipher Vigenere) Encode(input string) string {
	re := regexp.MustCompile("[^a-zA-Z]")
	input = re.ReplaceAllString(input, "")

	var chars = []rune(input)

	keyI := 0
	for i := range input {
		keyI = keyI % len(cipher.key)
		chars[i] = moveLetter(chars[i], int(cipher.key[keyI]-'a'))
		keyI++
	}

	return string(chars)
}

func (cipher Vigenere) Decode(input string) string {
	var chars = []rune(input)

	keyI := 0
	for i := range input {
		keyI = keyI % len(cipher.key)
		chars[i] = moveLetter(chars[i], -int(cipher.key[keyI]-'a'))
		keyI++
	}

	return string(chars)
}

func NewVigenere(key string) Cipher {
	if isKeyAffordable(key) {
		var cipher Cipher = &Vigenere{key}
		return cipher
	} else {
		return nil
	}
}

func isKeyAffordable(key string) bool {
	isRowOfA := true

	for _, c := range key {
		if !(c >= 'a' && c <= 'z') {
			return false
		}
		if c != 'a' {
			isRowOfA = false
		}
	}
	return !isRowOfA
}
