package form

import (
	"fmt"
	"io"

	"github.com/jzes/snakinha-go/board"
	"github.com/jzes/snakinha-go/fruit"
	"github.com/jzes/snakinha-go/snake"
)

type StdIOFormer struct {
	Board  board.Board
	Fruits fruit.Fruits
	Snake  snake.Snake
}

func (s *StdIOFormer) Form(w io.Writer, sprite string) {
	for y := 0; y < s.Board.SizeY; y++ {
		for x := 0; x < s.Board.SizeX; x++ {
			if s.Fruits.CheckFruits(x, y) { //substituir por um case e checar a cobrinha
				fmt.Fprint(w, "")
			} else {
				fmt.Fprint(w, sprite)
			}
		}
		fmt.Println("")
	}
}
