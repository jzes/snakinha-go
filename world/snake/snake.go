package snake

import (
	"github.com/jzes/snakinha-go/world/board"
	"github.com/jzes/snakinha-go/world/fruit"
)

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

// uma boa melhoria aqui é separar definitivamente a cabeça do corpo assim da pra usar o mesmo metodo que o checkSnake
func (s *Snake) AteItSelf() bool {
	head := s.Body[0]
	for _, bs := range s.Body[1:] {
		if bs.X == head.X && bs.Y == head.Y {
			return true
		}
	}
	return false
}

//uma melhoria aqui é não acessar diretamente o sizex e y do board e
// sim criar um metodo pra encapsular esse comportamento lá no board

func (s *Snake) ItsTheWall(b *board.Board) bool {
	head := s.Body[0]
	return head.X >= b.SizeX ||
		head.Y >= b.SizeY ||
		head.X < 0 ||
		head.Y < 0
}

func (s *Snake) Ate(f fruit.Fruit) bool {
	head := s.Body[0]
	if f.Equals(head.X, head.Y) {
		return true
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

func (s *Snake) Grow1() {
	lastBS := s.Body[len(s.Body)-1]
	s.Body = Body(append(s.Body, BodySection{
		X: lastBS.X + 1,
		Y: lastBS.Y}))
}
