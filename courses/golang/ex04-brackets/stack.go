package brackets

type Stack struct {
	arr []rune
}

func New() *Stack {
	var stack Stack = Stack{nil}
	return &stack
}

func (s *Stack) Push(val rune) {
	// s.arr[len(s.arr)-1] = val
	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() rune {
	if len(s.arr) == 0 {
		return 0
	}
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return val
}
