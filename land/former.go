package land

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/jzes/snakinha-go/world"
)

func Form(w io.Writer, world *world.World) {
	for y := 0; y < world.Board.SizeY; y++ {
		for x := 0; x < world.Board.SizeX; x++ {
			switch {
			case world.Fruit.Equals(x, y):
				fmt.Fprint(w, world.Fruit.Sprite)
			case world.Player.CheckSnake(x, y):
				fmt.Fprint(w, world.Player.Sprite)
			default:
				fmt.Fprint(w, world.Board.Sprite)
			}
		}
		fmt.Fprint(w, "|\n")
	}
	fmt.Fprint(w, strings.Repeat("-", world.Board.SizeX))
}

func Print(w io.Writer, screen tcell.Screen) error {
	land := strings.Split(w.(*bytes.Buffer).String(), "\n")

	for y, line := range land {
		for x, char := range line {
			screen.SetContent(x, y, rune(char), nil, tcell.StyleDefault)
		}
	}

	return nil
}
