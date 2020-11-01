package board

type Board struct {
	tiles [][]int
}

func NewBoard(tiles [][]int) *Board {
	return &Board{
		tiles: tiles,
	}
}

func (b *Board) ToString() string {
	return ""
}

func (b *Board) TileAt(row, col int) int {
	return b.tiles[row][col]
}

func (b *Board) Size() int {
	return len(b.tiles)
}

func (b *Board) Hamming() int {
	return 0
}

func (b *Board) Manhattan() int {
	return 0
}

func (b *Board) IsGoal() bool {
	return true
}

func (b *Board) Equals(y Board) bool {
	return true
}

//iterable

func (b *Board) IsSolvable() bool {
	invCount := getInvCounter(b.tiles, b.Size())
	return invCount%2 == 0
}

func getInvCounter(tiles [][]int, size int) int {
	inv_count := 0
	var i, j int
	for i = 0; i < size-1; i++ {
		for j = i + 1; j < size; j++ {
			if tiles[j][i] > 0 && tiles[j][i] > tiles[i][j] {
				inv_count++
			}
		}
	}
	return inv_count
}

func generateFinalState(size int) [][]int{
	var finalState [][]int
	counter := 1
	for i:= 0; i < size; i++ {
		for j := 0 < size; i++ {
			finalState[i][j] = counter
			counter++
		}
	}
	finalState[size-1][size-1] = 0
	return finalState
}