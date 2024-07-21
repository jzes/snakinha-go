package main

import (
	"os"

	"github.com/jzes/snakinha-go/board"
	"github.com/jzes/snakinha-go/form"
)

func main() {
	b := board.Board{
		SizeY: 10,
		SizeX: 30,
	}
	f := form.StdIOFormer{
		Board: b,
	}
	f.Form(os.Stdout, "x")
}
