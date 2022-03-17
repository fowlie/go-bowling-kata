package main

type spare struct {
	current roll
	next    *frame
}

func (s *spare) pins() (int, int) {
	return s.current.first, s.current.second
}

func (s *spare) score() int {
	next, _ := (*s.next).pins()
	return 10 + next
}
