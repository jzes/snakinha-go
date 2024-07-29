package snake

type Snake struct {
	Body   Body
	Sprite string
}

func New(b Body, s string) Snake {
	return Snake{b, s}
}

func (s *Snake) CheckSnake(x, y int) bool {
	for _, bP := range s.Body {
		if bP.X == x && bP.Y == y {
			return true
		}
	}
	return false
}

func (s *Snake) ToDown() {
	head := s.Body[0]
	body := s.Body[:len(s.Body)-1]

	newHead := []BodySection{{X: head.X, Y: head.Y + 1}}
	s.Body = Body(append(newHead, body...))
}

func (s *Snake) ToUp() {
	head := s.Body[0]
	body := s.Body[:len(s.Body)-1]

	newHead := []BodySection{{X: head.X, Y: head.Y - 1}}
	s.Body = Body(append(newHead, body...))
}

func (s *Snake) ToRight() {
	head := s.Body[0]
	body := s.Body[:len(s.Body)-1]

	newHead := []BodySection{{X: head.X + 1, Y: head.Y}}
	s.Body = Body(append(newHead, body...))
}

func (s *Snake) ToLeft() {
	head := s.Body[0]
	body := s.Body[:len(s.Body)-1]

	newHead := []BodySection{{X: head.X - 1, Y: head.Y}}
	s.Body = Body(append(newHead, body...))
}
