package board

type Board struct {
	tiles [][]int
	index int
}

func NewBoard(tiles [][]int, index int) *Board {
	return &Board{
		tiles: tiles,
		index: index,
	}
}

func (b *Board) Size() int {
	return len(b.tiles)
}

func (b *Board) Manhattan(finalState [][]int) int {
	cost := 0
	rowSize := b.Size()

	for i := 0; i < rowSize; i++ {
		for j := 0; j < rowSize; j++ {
			if b.tiles[i][j] == 0 {
				continue
			}
			x, y := findIndex(b.tiles[i][j], finalState)
			cost = cost + abs(i-x) + abs(j-y)
		}
	}
	return cost
}

func (b *Board) Equals(y Board) bool {
	size := b.Size()
	sizeOther := y.Size()
	if size != sizeOther {
		return false
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b.tiles[i][j] != y.tiles[i][j] {
				return false
			}
		}
	}
	return true
}

func (b *Board) Neightbours() ([]Board, []string) {
	var result []Board
	var path []string
	var x, y int
	for i := 0; i < b.Size(); i++ {
		for j := 0; j < b.Size(); j++ {
			if b.tiles[i][j] == 0 {
				x, y = i, j
			}
		}
	}

	if x-1 >= 0 {
		board := copyBoard(*b)
		board.tiles[x][y], board.tiles[x-1][y] = board.tiles[x-1][y], board.tiles[x][y]
		result = append(result, board)
		path = append(path, "down")
	}
	if x+1 < b.Size() {
		board := copyBoard(*b)
		board.tiles[x][y], board.tiles[x+1][y] = board.tiles[x+1][y], board.tiles[x][y]
		result = append(result, board)
		path = append(path, "up")
	}
	if y-1 >= 0 {

		board := copyBoard(*b)
		board.tiles[x][y], board.tiles[x][y-1] = board.tiles[x][y-1], board.tiles[x][y]
		result = append(result, board)
		path = append(path, "right")
	}
	if y+1 < b.Size() {
		board := copyBoard(*b)
		board.tiles[x][y], board.tiles[x][y+1] = board.tiles[x][y+1], board.tiles[x][y]
		result = append(result, board)
		path = append(path, "left")
	}
	return result, path
}

func (b *Board) IsSolvable() bool {
	invCount := getInvCounter(b.tiles, b.Size())
	return invCount%2 == 0
}

func copyBoard(b Board) Board {
	tiles := make([][]int, len(b.tiles))
	for i := range b.tiles {
		tiles[i] = make([]int, len(b.tiles[i]))
		copy(tiles[i], b.tiles[i])
	}
	return Board{tiles: tiles}
}

func getInvCounter(tiles [][]int, size int) int {
	invCount := 0
	row := toRow(tiles)
	var i, j int
	for i = 0; i < size*size-1; i++ {
		for j = i + 1; j < size*size; j++ {
			if row[i] != 0 && row[j] != 0 && row[i] > row[j] {
				invCount++
			}
		}
	}
	return invCount
}

func toRow(arr [][]int) []int {
	var row []int
	for i, _ := range arr {
		for j, _ := range arr {
			row = append(row, arr[i][j])
		}
	}
	return row
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GenerateFinalState(size, index int) [][]int {
	var finalState [][]int
	var i, j int
	counter := 0
	elem := 0
	if index < 0 {
		index = size*size - 1
	}
	for i = 0; i < size; i++ {
		var row []int
		for j = 0; j < size; j++ {
			if index == counter {
				row = append(row, 0)
				counter++
			} else {
				elem++
				row = append(row, elem)
				counter++
			}
		}
		finalState = append(finalState, row)
	}

	return finalState
}

func findIndex(elem int, tiles [][]int) (int, int) {
	size := len(tiles)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if elem == tiles[i][j] {
				return i, j
			}
		}
	}
	return -1, -1
}
