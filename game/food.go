package game

import "math/rand"

type Food struct {
	Position
}

func NewFood(board *Board, snake *Snake) *Food {
	f := &Food{}
	f.Spawn(board, snake)
	return f
}

func (f *Food) Spawn(board *Board, snake *Snake) {
	occupied := make(map[Position]bool)
	n := len(snake.Path)
	for i := 0; i < snake.SegLen && n-1-i >= 0; i++ {
		occupied[snake.Path[n-1-i]] = true
	}

	var empty []Position
	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			pos := Position{X: x, Y: y}
			if !occupied[pos] {
				empty = append(empty, pos)
			}
		}
	}

	if len(empty) > 0 {
		f.Position = empty[rand.Intn(len(empty))]
	}
}
