package main

import "testing"

func Test_GutterBalls(t *testing.T) {
	g := NewBowlingGame()
	rollMany(g, 10, 0, 0)
	assertEquals(t, 0, g.GetTotalScore())
}

func Test_Threes(t *testing.T) {
	g := NewBowlingGame()
	rollMany(g, 10, 3, 3)
	assertEquals(t, 60, g.GetTotalScore())
}

func Test_Spare(t *testing.T) {
	g := NewBowlingGame()
	g.Roll(4, 6)
	g.Roll(3, 5)
	rollMany(g, 8, 0, 0)
	assertEquals(t, 21, g.GetTotalScore())
}

func Test_Strike(t *testing.T) {
	g := NewBowlingGame()
	g.Roll(10, 0)
	g.Roll(5, 3)
	rollMany(g, 8, 0, 0)
	assertEquals(t, 26, g.GetTotalScore())
}

func rollMany(b *BowlingGame, count int, first int, second int) {
	for i := 0; i < count; i++ {
		b.Roll(first, second)
	}
}

func assertEquals(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Fatalf("Expected %d score but was %d", expected, actual)
	}
}
