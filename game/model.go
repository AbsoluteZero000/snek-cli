package game

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	DefaultWidth     = 20
	DefaultHeight    = 20
	InitialTickSpeed = 200 * time.Millisecond
	MinTickSpeed     = 60 * time.Millisecond
	SpeedDecrement   = 8 * time.Millisecond
	FrameDuration    = 16 * time.Millisecond
)

type GameState int

const (
	StatePlaying GameState = iota
	StateGameOver
)

type tickMsg struct{}
type frameMsg struct{}

type GameModel struct {
	board        *Board
	styles       *Styles
	snake        *Snake
	food         *Food
	state        GameState
	score        int
	tickInterval time.Duration
	width        int
	height       int
	ready        bool
}

func New() *GameModel {
	m := &GameModel{
		board:        NewBoard(DefaultWidth, DefaultHeight),
		styles:       NewStyles(),
		snake:        NewSnake(DefaultWidth, DefaultHeight),
		state:        StatePlaying,
		tickInterval: InitialTickSpeed,
	}
	m.food = NewFood(m.board, m.snake)
	return m
}

func (m *GameModel) tick() tea.Cmd {
	return tea.Tick(m.tickInterval, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m *GameModel) frameTick() tea.Cmd {
	return tea.Tick(FrameDuration, func(t time.Time) tea.Msg {
		return frameMsg{}
	})
}

func (m *GameModel) Init() tea.Cmd {
	return tea.Batch(m.tick(), m.frameTick())
}
