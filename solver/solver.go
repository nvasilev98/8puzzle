package solver

import (
	"fmt"

	"github.com/nvasilev98/8puzzle/board"
)

func Solve(board board.Board, path []string, level int, states []board.Board, threshold int, finalState [][]int) bool {
	var solve bool
	distance := board.Manhattan(finalState)
	if distance == 0 {
		fmt.Println(len(path))
		fmt.Println(path)
		return true
	}
	if contains(states, board) {
		return false
	}
	states = append(states, board)
	if level+distance > threshold {
		if len(states) > 0 {
			states = states[:len(states)-1]
		}
		return false
	}
	nb, nbpath := board.Neightbours()

	for i, _ := range nb {
		path = append(path, nbpath[i])
		solve = Solve(nb[i], path, level+1, states, threshold, finalState)
		if solve {
			return solve
		}
		path = path[:len(path)-1]
	}
	return solve

}

func contains(s []board.Board, e board.Board) bool {
	for _, a := range s {
		if a.Equals(e) {
			return true
		}
	}
	return false
}
