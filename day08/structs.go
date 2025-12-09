package main

type JunctionBox struct {
	X int
	Y int
	Z int
}

type Circuit struct {
	junctionBox []JunctionBox
}

type ClosestJunctionBoxes struct {
	jBox1    JunctionBox
	jBox2    JunctionBox
	distance float64
}
