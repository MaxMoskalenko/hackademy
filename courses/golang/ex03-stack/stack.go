package stack

type Stack struct {
	arr []int
}

func New() *Stack {
	var stack Stack = Stack{nil}
	return &stack
}

func (s *Stack) Push(val int) {
	// s.arr[len(s.arr)-1] = val
	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() int {
	if len(s.arr) == 0 {
		return 0
	}
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return val
}
