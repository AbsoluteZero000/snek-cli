package game

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *GameModel) View() string {
	if !m.ready {
		return "Loading..."
	}

	var b strings.Builder
	for y := 0; y < m.board.Height; y++ {
		for x := 0; x < m.board.Width; x++ {
			cell := string(m.board.Grid[y][x])
			b.WriteString(m.styles.Cell.Render(cell))
		}
		if y < m.board.Height-1 {
			b.WriteString("\n")
		}
	}

	board := m.styles.Board.Render(b.String())
	score := m.styles.Score.Render("Score: 0")

	return lipgloss.JoinVertical(lipgloss.Center, board, score)
}
