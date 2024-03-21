package deque

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return value, true
}

func (s *Stack) Peak() int {
	return (*s)[len(*s)-1]
}
