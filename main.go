package main

import (
	"fmt"
	"os"

	"github.com/absolutezero000/snek-cli/game"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := game.New()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
