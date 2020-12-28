package helpers

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

type IntStack []int

func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *IntStack) Push(v int) {
	*s = append(*s, v)
}

func (s *IntStack) Pop() int {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}
