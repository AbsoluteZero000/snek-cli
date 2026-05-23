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

	case tickMsg:
		if m.state != StatePlaying {
			return m, nil
		}

		m.snake.Move()

		if !m.board.InBounds(m.snake.Head().X, m.snake.Head().Y) {
			m.state = StateGameOver
			return m, nil
		}

		if m.snake.CollidesWithSelf() {
			m.state = StateGameOver
			return m, nil
		}

		return m, m.tick()
	}

	return m, nil
}
