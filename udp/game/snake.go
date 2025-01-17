package game

type Snake struct {
	body      []Point
	direction Point
}

func (s *Snake) Body() []Point {
	return s.body
}

func (s *Snake) SetBody(body []Point) {
	s.body = body
}

func (s *Snake) Direction() Point {
	return s.direction
}

func (s *Snake) SetDirection(direction Point) {
	s.direction = direction
}

func NewSnake() *Snake {
	return &Snake{
		body: []Point{{
			X: 0,
			Y: Setting.Height / 2,
		}},
		direction: Point{
			X: 1,
			Y: 0,
		},
	}
}

func (s *Snake) Move() {
	head := s.Body()[0]
	newX := head.X + s.Direction().X
	newY := head.Y + s.Direction().Y

	if newX >= Setting.Width && s.Direction().X == 1 {
		newX = 0
	}

	if newX == 0 && s.Direction().X == -1 {
		newX = Setting.Width - 1
	}

	if newY >= Setting.Height && s.Direction().Y == 1 {
		newY = 0
	}

	if newY == 0 && s.Direction().Y == -1 {
		newY = Setting.Height - 1
	}

	s.SetBody(append([]Point{{X: newX, Y: newY}}, s.Body()[:len(s.Body())-1]...))
}

func (s *Snake) ChangeDirection(newDir Point) {
	if newDir.X != 0 && s.Direction().X != 0 {
		return
	}

	if newDir.Y != 0 && s.Direction().Y != 0 {
		return
	}

	s.SetDirection(newDir)
}

func (s *Snake) Grow() {
	tail := s.Body()[len(s.Body())-1]
	s.SetBody(append(s.Body(), tail))
}
