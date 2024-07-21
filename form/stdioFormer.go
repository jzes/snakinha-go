package form

import (
	"fmt"
	"io"

	"github.com/jzes/snakinha-go/board"
)

type StdIOFormer struct {
	Board board.Board
}

func (s *StdIOFormer) Form(w io.Writer, sprite string) {
	for y := 0; y < s.Board.SizeY; y++ {
		for x := 0; x < s.Board.SizeX; x++ {
			fmt.Fprint(w, sprite)
		}
		fmt.Println("")
	}
}
