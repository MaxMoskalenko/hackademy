package cipher

import "regexp"

func NewCaesar() Cipher {
	return NewShift(3)
}

type Shift struct {
	num int
}

func NewShift(num int) Cipher {
	if (num >= 1 && num <= 25) || (num <= -1 && num >= -25) {
		var cipher Cipher = &Shift{num}
		return cipher
	} else {
		return nil
	}
}

func (cipher Shift) Encode(input string) string {

	re := regexp.MustCompile("[^a-zA-Z]")
	input = re.ReplaceAllString(input, "")

	var chars = []rune(input)

	for i := range input {
		if !checkIsLetter(chars[i]) {
			chars = append(chars[:i], chars[i+1:]...)
		} else {
			chars[i] = moveLetter(chars[i], cipher.num)
		}
	}
	return string(chars)
}

func (cipher Shift) Decode(input string) string {
	var chars = []rune(input)
	for i := range input {

		chars[i] = moveLetter(chars[i], -cipher.num)
	}
	return string(chars)
}

func moveLetter(c rune, num int) rune {
	if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}
	if c+rune(num) > 'z' {
		return c + rune(num) - 'z' + 'a' - 1
	}
	if c+rune(num) < 'a' {
		return 'z' + c + rune(num) - 'a' + 1
	}

	return c + rune(num)
}

func checkIsLetter(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}
