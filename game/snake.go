package game

type Position struct {
	X, Y int
}

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

type Snake struct {
	Body    []Position
	Dir     Direction
	NextDir Direction
	growing bool
}

func NewSnake(width, height int) *Snake {
	cx, cy := width/2, height/2
	return &Snake{
		Body: []Position{
			{X: cx, Y: cy},
			{X: cx - 1, Y: cy},
			{X: cx - 2, Y: cy},
		},
		Dir:     DirRight,
		NextDir: DirRight,
	}
}

func (s *Snake) Head() Position {
	return s.Body[0]
}

func (s *Snake) Move() {
	s.Dir = s.NextDir
	head := s.Head()
	var newHead Position
	switch s.Dir {
	case DirUp:
		newHead = Position{X: head.X, Y: head.Y - 1}
	case DirDown:
		newHead = Position{X: head.X, Y: head.Y + 1}
	case DirLeft:
		newHead = Position{X: head.X - 1, Y: head.Y}
	case DirRight:
		newHead = Position{X: head.X + 1, Y: head.Y}
	}

	if s.growing {
		s.Body = append([]Position{newHead}, s.Body...)
		s.growing = false
	} else {
		s.Body = append([]Position{newHead}, s.Body[:len(s.Body)-1]...)
	}
}

func (s *Snake) ChangeDir(dir Direction) {
	if s.Dir == DirUp && dir == DirDown {
		return
	}
	if s.Dir == DirDown && dir == DirUp {
		return
	}
	if s.Dir == DirLeft && dir == DirRight {
		return
	}
	if s.Dir == DirRight && dir == DirLeft {
		return
	}
	s.NextDir = dir
}

func (s *Snake) CollidesWithSelf() bool {
	head := s.Head()
	for i := 1; i < len(s.Body); i++ {
		if s.Body[i] == head {
			return true
		}
	}
	return false
}

func (s *Snake) StartGrow() {
	s.growing = true
}
