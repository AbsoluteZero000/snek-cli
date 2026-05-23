package game

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Board lipgloss.Style
	Cell  lipgloss.Style
	Score lipgloss.Style
}

func NewStyles() *Styles {
	return &Styles{
		Board: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1),
		Cell: lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center),
		Score: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")),
	}
}
