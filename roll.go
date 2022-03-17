package main

type roll struct {
	first, second int
}

func (r *roll) pins() (int, int) {
	return r.first, r.second
}

func (r *roll) score() int {
	return sumOf(r.pins())
}
