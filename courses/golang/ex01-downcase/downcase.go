package downcase

func Downcase(input string) (string, error) {
	var chars = []rune(input)
	for i, c := range input {
		if c <= 'Z' && c >= 'A' {
			chars[i] = c - 'A' + 'a'
		}
	}
	return string(chars), nil
}
