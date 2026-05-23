package game

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *GameModel) View() string {
	if !m.ready {
		return "Loading..."
	}

	m.drawGrid()

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
	score := m.styles.Score.Render(m.scoreString())

	if m.state == StateGameOver {
		over := m.styles.Board.Render("  GAME OVER  \n\n  Press q to quit")
		return lipgloss.JoinVertical(lipgloss.Center, over, score)
	}

	return lipgloss.JoinVertical(lipgloss.Center, board, score)
}

func (m *GameModel) scoreString() string {
	return "Score: 0"
}

func (m *GameModel) drawGrid() {
	for y := 0; y < m.board.Height; y++ {
		for x := 0; x < m.board.Width; x++ {
			m.board.Grid[y][x] = ' '
		}
	}

	for _, p := range m.snake.Body {
		m.board.Grid[p.Y][p.X] = '█'
	}
}
