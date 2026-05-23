package game

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Board    lipgloss.Style
	Cell     lipgloss.Style
	Food     lipgloss.Style
	Score    lipgloss.Style
	GameOver lipgloss.Style
}

func NewStyles() *Styles {
	return &Styles{
		Board: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1),
		Cell: lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#00FF00")),
		Food: lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#FF4040")),
		Score: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")),
		GameOver: lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			Padding(1, 3).
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")).
			BorderForeground(lipgloss.Color("#FF0000")),
	}
}
