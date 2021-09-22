package brackets

func Bracket(input string) (bool, error) {
	var s Stack = *New()

	for _, c := range input {
		if c == '}' {
			if s.Pop() != '{' {
				return false, nil
			}
		}
		if c == ')' {
			if s.Pop() != '(' {
				return false, nil
			}
		}
		if c == ']' {
			if s.Pop() != '[' {
				return false, nil
			}
		}

		if c == '[' || c == '{' || c == '(' {
			s.Push(c)
		}

	}
	if len(s.arr) != 0 {
		return false, nil
	}
	return true, nil

}
