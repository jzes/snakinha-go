package main

import (
	"bytes"
	"log"

	"github.com/gdamore/tcell"
	"github.com/jzes/snakinha-go/board"
	"github.com/jzes/snakinha-go/fruit"
	"github.com/jzes/snakinha-go/land"
	"github.com/jzes/snakinha-go/snake"
)

func main() {
	board1 := board.New(" ", 10, 20)

	gardem := fruit.NewGardem("*")
	gardem.WithFruit(fruit.New(2, 2))
	gardem.WithFruit(fruit.New(7, 8))

	body := snake.Body{}
	body.WithSection(snake.BodySection{X: 9, Y: 1})
	body.WithSection(snake.BodySection{X: 10, Y: 1})
	body.WithSection(snake.BodySection{X: 11, Y: 1})
	body.WithSection(snake.BodySection{X: 12, Y: 1})
	player1 := snake.New(body, "#")

	terraFormer := land.Former{}
	terraFormer.WithBoard(board1)
	terraFormer.WithGardem(gardem)
	terraFormer.WithSnake(player1)

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("error on creating a new screen for tcell lib")
	}
	err = screen.Init()
	if err != nil {
		log.Fatal("error on init the screen for tcell lib")
	}
	screen.SetStyle(tcell.StyleDefault)

	tcellFmt := land.TCellFmt{
		Screen: screen,
		Style:  tcell.StyleDefault,
	}

	buff := bytes.Buffer{}
	terraFormer.Form(&buff)

	tcellFmt.Print(&buff)

	for {
		screen.Show()

		event := screen.PollEvent()

		switch event := event.(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				return
			}
			if event.Key() == tcell.KeyDown {
				player1.ToDown()

				terraFormer.WithSnake(player1)

				buff.Reset()
				terraFormer.Form(&buff)
				tcellFmt.Print(&buff)
			}
			if event.Key() == tcell.KeyUp {
				player1.ToUp()

				terraFormer.WithSnake(player1)

				buff.Reset()
				terraFormer.Form(&buff)
				tcellFmt.Print(&buff)
			}
			if event.Key() == tcell.KeyRight {
				player1.ToRight()

				terraFormer.WithSnake(player1)

				buff.Reset()
				terraFormer.Form(&buff)
				tcellFmt.Print(&buff)
			}
			if event.Key() == tcell.KeyLeft {
				player1.ToLeft()

				terraFormer.WithSnake(player1)

				buff.Reset()
				terraFormer.Form(&buff)
				tcellFmt.Print(&buff)
			}
		}
	}

}
