package main

import "fmt"

type Nonogram struct {
	rows [][]int
	cols [][]int
	grid [][]int
}

func getOneNono() Nonogram {
	rows := [][]int{
		{3, 1},
		{1, 2},
		{2, 1},
		{1, 1},
		{1},
	}
	cols := [][]int{
		{3},
		{1, 2},
		{1, 1},
		{1, 1},
		{3},
	}
	grid := [][]int{
		{1, 1, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
	}
	return Nonogram{rows, cols, grid}
}

func main() {
	fmt.Printf("1")
}
