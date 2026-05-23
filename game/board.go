package game

type Board struct {
	Width  int
	Height int
	Grid   [][]rune
}

func NewBoard(width, height int) *Board {
	grid := make([][]rune, height)
	for y := range grid {
		grid[y] = make([]rune, width)
		for x := range grid[y] {
			grid[y][x] = ' '
		}
	}
	return &Board{
		Width:  width,
		Height: height,
		Grid:   grid,
	}
}

func (b *Board) InBounds(x, y int) bool {
	return x >= 0 && x < b.Width && y >= 0 && y < b.Height
}
