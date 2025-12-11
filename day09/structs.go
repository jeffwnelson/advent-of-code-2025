package main

type Point struct {
	X int
	Y int
}

type AreaResult struct {
	P1   Point
	P2   Point
	Area int
}

type Tile struct {
	Position Point
	Type     rune // 'R' = Red, 'G' = Green
}
