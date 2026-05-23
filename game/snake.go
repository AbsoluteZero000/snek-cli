package game

type Position struct {
	X, Y int
}

type FloatPosition struct {
	X, Y float64
}

type Direction int

const (
	DirUp Direction = iota
	DirDown
	DirLeft
	DirRight
)

type Snake struct {
	Path     []Position
	SegLen   int
	Dir      Direction
	NextDir  Direction
	growing  bool
	Progress float64
}

func NewSnake(width, height int) *Snake {
	cx, cy := width/2, height/2
	return &Snake{
		Path: []Position{
			{X: cx - 2, Y: cy},
			{X: cx - 1, Y: cy},
			{X: cx, Y: cy},
		},
		SegLen:  3,
		Dir:     DirRight,
		NextDir: DirRight,
	}
}

func (s *Snake) Head() Position {
	return s.Path[len(s.Path)-1]
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

	s.Path = append(s.Path, newHead)
	if s.growing {
		s.SegLen++
		s.growing = false
	}

	if len(s.Path) > s.SegLen*3 {
		s.Path = s.Path[len(s.Path)-s.SegLen*3:]
	}
}

func (s *Snake) InterpolatedPositions() []FloatPosition {
	n := len(s.Path)
	if n == 0 || s.SegLen == 0 {
		return nil
	}

	positions := make([]FloatPosition, s.SegLen)
	for i := 0; i < s.SegLen; i++ {
		curr := s.Path[n-1-i]

		if n-2-i >= 0 {
			prev := s.Path[n-2-i]
			positions[i] = FloatPosition{
				X: float64(prev.X) + float64(curr.X-prev.X)*s.Progress,
				Y: float64(prev.Y) + float64(curr.Y-prev.Y)*s.Progress,
			}
		} else {
			positions[i] = FloatPosition{X: float64(curr.X), Y: float64(curr.Y)}
		}
	}
	return positions
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
	n := len(s.Path)
	if n < 2 {
		return false
	}
	head := s.Path[n-1]
	limit := 2
	if n-limit < 0 {
		limit = n
	}
	for i := limit; i <= s.SegLen && n-1-i >= 0; i++ {
		if s.Path[n-1-i] == head {
			return true
		}
	}
	return false
}

func (s *Snake) StartGrow() {
	s.growing = true
}
