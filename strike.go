package main

type strike struct {
	next *frame
}

func (s *strike) pins() (int, int) {
	return 10, 0
}

func (s *strike) score() int {
	return 10 + sumOf((*s.next).pins())
}
