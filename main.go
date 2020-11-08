package main

import (
	"fmt"
	"math"

	"github.com/nvasilev98/8puzzle/board"
	"github.com/nvasilev98/8puzzle/solver"
)

func insertMatrix(size int) [][]int {

	var tiles [][]int
	for i := 0; i < size; i++ {
		tmp := make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Printf("Input number [%d][%d]: ", i, j)
			fmt.Scanf("%d", &tmp[j])
		}
		tiles = append(tiles, tmp)
	}
	return tiles
}

func main() {
	fmt.Print("Enter size: ")
	var size float64
	var index int
	fmt.Scanln(&size)
	fmt.Print("Enter index: ")
	fmt.Scanln(&index)
	sqrtSize := math.Sqrt(size + 1)
	tiles := insertMatrix(int(sqrtSize))
	bord := board.NewBoard(tiles, index)
	res := bord.IsSolvable()
	if res == false {
		fmt.Println("not solvable")
		return
	}
	finalStale := board.GenerateFinalState(int(sqrtSize), index)
	var states []board.Board
	startDistance := bord.Manhattan(finalStale)
	var path []string
	solve := false
	for !solve {
		solve = solver.Solve(*bord, path, 0, states, startDistance, finalStale)
		states = nil
		startDistance++
	}
}
