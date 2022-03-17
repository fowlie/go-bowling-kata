package main

type BowlingGame struct {
	throws []frame
	index  int
}

func NewBowlingGame() *BowlingGame {
	return &BowlingGame{
		throws: make([]frame, 10),
		index:  0,
	}
}

func (g *BowlingGame) Roll(first int, second int) {
	var f frame
	if first+second < 10 {
		f = &roll{first, second}
	} else if first == 10 {
		f = &strike{&g.throws[g.index+1]}
	} else {
		f = &spare{roll{first, second}, &g.throws[g.index+1]}
	}
	g.add(f)
}

func (g *BowlingGame) GetTotalScore() int {
	total := 0
	for _, frame := range g.throws {
		total += frame.score()
	}
	return total
}

func (g *BowlingGame) add(f frame) {
	g.throws[g.index] = f
	g.index++
}

func sumOf(a int, b int) int {
	return a + b
}
