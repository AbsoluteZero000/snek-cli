package game

import (
	"math"
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

type cellRender struct {
	x, y int
	ch   rune
}

func (m *GameModel) drawGrid() {
	for y := 0; y < m.board.Height; y++ {
		for x := 0; x < m.board.Width; x++ {
			m.board.Grid[y][x] = ' '
		}
	}

	positions := m.snake.InterpolatedPositions()
	var renders []cellRender

	for i, fp := range positions {
		cells := floatToCells(fp.X, fp.Y)
		for _, c := range cells {
			if c.x >= 0 && c.x < m.board.Width && c.y >= 0 && c.y < m.board.Height {
				if i == 0 {
					renders = append(renders, c)
				} else {
					renders = append(renders, c)
				}
			}
		}
	}

	for _, r := range renders {
		m.board.Grid[r.y][r.x] = r.ch
	}
}

func floatToCells(fx, fy float64) []cellRender {
	ix := int(math.Floor(fx))
	iy := int(math.Floor(fy))
	fx -= float64(ix)
	fy -= float64(iy)

	var cells []cellRender

	vert := fy > 0.001
	horz := fx > 0.001

	if !vert && !horz {
		return []cellRender{{ix, iy, '█'}}
	}

	if vert {
		topOverlap := 1.0 - fy
		if topOverlap >= 0.7 {
			cells = append(cells, cellRender{ix, iy, '█'})
		} else if topOverlap >= 0.3 {
			cells = append(cells, cellRender{ix, iy, '▄'})
		}

		if fy >= 0.7 {
			cells = append(cells, cellRender{ix, iy + 1, '█'})
		} else if fy >= 0.3 {
			cells = append(cells, cellRender{ix, iy + 1, '▀'})
		}

		return cells
	}

	if horz {
		leftOverlap := 1.0 - fx
		if leftOverlap >= 0.7 {
			cells = append(cells, cellRender{ix, iy, '█'})
		} else if leftOverlap >= 0.3 {
			cells = append(cells, cellRender{ix, iy, '▐'})
		}

		if fx >= 0.7 {
			cells = append(cells, cellRender{ix + 1, iy, '█'})
		} else if fx >= 0.3 {
			cells = append(cells, cellRender{ix + 1, iy, '▌'})
		}

		return cells
	}

	return cells
}
