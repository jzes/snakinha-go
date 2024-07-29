package land

import (
	"fmt"
	"io"
	"strings"

	"github.com/jzes/snakinha-go/board"
	"github.com/jzes/snakinha-go/fruit"
	"github.com/jzes/snakinha-go/snake"
)

type Former struct {
	board  board.Board
	gardem fruit.Gardem
	snake  snake.Snake
}

func (f *Former) WithBoard(b board.Board) {
	f.board = b
}

func (f *Former) WithGardem(g fruit.Gardem) {
	f.gardem = g
}

func (f *Former) WithSnake(s snake.Snake) {
	f.snake = s
}

func (s *Former) Form(w io.Writer) {
	for y := 0; y < s.board.SizeY; y++ {
		for x := 0; x < s.board.SizeX; x++ {
			switch {
			case s.gardem.Fruits.CheckFruits(x, y):
				fmt.Fprint(w, s.gardem.Sprite)
			case s.snake.CheckSnake(x, y):
				fmt.Fprint(w, s.snake.Sprite)
			default:
				fmt.Fprint(w, s.board.Sprite)
			}
		}
		fmt.Fprint(w, "|\n")
	}
	fmt.Fprint(w, strings.Repeat("-", s.board.SizeX))
}
