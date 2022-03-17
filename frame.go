package main

type frame interface {
	pins() (int, int)
	score() int
}
