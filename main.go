package main

import (
	"os"

	"github.com/jzes/snakinha-go/board"
	"github.com/jzes/snakinha-go/form"
	"github.com/jzes/snakinha-go/fruit"
)

func main() {
	b := board.Board{
		SizeY: 10,
		SizeX: 30,
	}
	fs := fruit.Fruits{
		{X: 2, Y: 2},
		{X: 3, Y: 4},
	}
	fr := form.StdIOFormer{
		Board:  b,
		Fruits: fs,
	}

	fr.Form(os.Stdout, "")

}
