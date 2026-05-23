package game

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	DefaultWidth  = 20
	DefaultHeight = 20
)

type GameModel struct {
	board  *Board
	styles *Styles
	score  int
	width  int
	height int
	ready  bool
}

func New() *GameModel {
	return &GameModel{
		board:  NewBoard(DefaultWidth, DefaultHeight),
		styles: NewStyles(),
		score:  0,
	}
}

func (m *GameModel) Init() tea.Cmd {
	return nil
}
