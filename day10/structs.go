package main

type Machine struct {
	Lights  string
	Buttons [][]int
	Joltage []int
}

type Solution struct {
	Buttons []int // location of buttons pressed
	Count   int   // number of buttons pressed
}
