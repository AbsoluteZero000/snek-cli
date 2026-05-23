package game

import (
	"fmt"
	"math"
	"strings"
	"time"

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
			ch := m.board.Grid[y][x]
			var styled string
			if ch == '●' {
				styled = m.styles.Food.Render(string(ch))
			} else {
				styled = m.styles.Cell.Render(string(ch))
			}
			b.WriteString(styled)
		}
		if y < m.board.Height-1 {
			b.WriteString("\n")
		}
	}

	score := m.styles.Score.Render(m.scoreString())

	if m.state == StateGameOver {
		board := m.styles.Board.Render(b.String())
		over := m.styles.GameOver.Render(
			"GAME OVER\n\n" +
				fmt.Sprintf("Score: %d\n", m.score) +
				"Speed: " + speedLabel(m.tickInterval) + "\n\n" +
				"r  restart\n" +
				"q  quit",
		)
		return lipgloss.JoinVertical(lipgloss.Center, board, over, score)
	}

	board := m.styles.Board.Render(b.String())
	return lipgloss.JoinVertical(lipgloss.Center, board, score)
}

func (m *GameModel) scoreString() string {
	return fmt.Sprintf("Score: %d  |  Speed: %s", m.score, speedLabel(m.tickInterval))
}

func speedLabel(d time.Duration) string {
	s := int(InitialTickSpeed / d)
	return fmt.Sprintf("%dx", s)
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

	m.board.Grid[m.food.Y][m.food.X] = '●'

	positions := m.snake.InterpolatedPositions()
	var renders []cellRender

	for _, fp := range positions {
		cells := floatToCells(fp.X, fp.Y)
		renders = append(renders, cells...)
	}

	for _, r := range renders {
		if r.x >= 0 && r.x < m.board.Width && r.y >= 0 && r.y < m.board.Height {
			m.board.Grid[r.y][r.x] = r.ch
		}
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
