package game

import tea "github.com/charmbracelet/bubbletea"

func (m *GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "w", "k":
			m.snake.ChangeDir(DirUp)
		case "down", "s", "j":
			m.snake.ChangeDir(DirDown)
		case "left", "a", "h":
			m.snake.ChangeDir(DirLeft)
		case "right", "d", "l":
			m.snake.ChangeDir(DirRight)
		}
		return m, nil

	case frameMsg:
		if m.state == StatePlaying {
			m.snake.Progress += float64(FrameDuration) / float64(m.tickInterval)
			if m.snake.Progress > 1.0 {
				m.snake.Progress = 1.0
			}
		}
		return m, m.frameTick()

	case tickMsg:
		if m.state != StatePlaying {
			return m, nil
		}

		m.snake.Move()
		m.snake.Progress = 0

		if !m.board.InBounds(m.snake.Head().X, m.snake.Head().Y) {
			m.state = StateGameOver
			return m, nil
		}

		if m.snake.CollidesWithSelf() {
			m.state = StateGameOver
			return m, nil
		}

		if m.snake.Head() == m.food.Position {
			m.snake.StartGrow()
			m.score++
			m.food.Spawn(m.board, m.snake)
			if m.tickInterval > MinTickSpeed {
				m.tickInterval -= SpeedDecrement
			}
		}

		return m, m.tick()
	}

	return m, nil
}
