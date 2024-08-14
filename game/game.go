package game

import (
	"bytes"

	"github.com/gdamore/tcell"
	"github.com/jzes/snakinha-go/land"
	"github.com/jzes/snakinha-go/snake"
)

type Game struct {
	screen        tcell.Screen
	player        snake.Snake
	terraFormer   land.Former
	terraFormater land.TCellFmt
	buff          bytes.Buffer
}

func New() *Game {
	return &Game{}
}

func (g *Game) WithScreen(s tcell.Screen) *Game {
	g.screen = s
	return g
}

func (g *Game) WithPlayer(p snake.Snake) *Game {
	g.player = p
	return g
}

func (g *Game) WithFormer(tf land.Former) *Game {
	g.terraFormer = tf
	return g
}

func (g *Game) WithFormater(tfmt land.TCellFmt) *Game {
	g.terraFormater = tfmt
	return g
}

func (g *Game) WithBuffer(b bytes.Buffer) *Game {
	g.buff = b
	return g
}

// apllicar um pouco de dry nesses metodos de tecla
func (g Game) Loop() {

	for {
		g.screen.Show()

		event := g.screen.PollEvent()

		switch event := event.(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				return
			}
			if event.Key() == tcell.KeyDown {
				g.player.ToDown()

				g.terraFormer.WithSnake(g.player)

				g.buff.Reset()
				g.terraFormer.Form(&g.buff)
				g.terraFormater.Print(&g.buff)
			}
			if event.Key() == tcell.KeyUp {
				g.player.ToUp()

				g.terraFormer.WithSnake(g.player)

				g.buff.Reset()
				g.terraFormer.Form(&g.buff)
				g.terraFormater.Print(&g.buff)
			}
			if event.Key() == tcell.KeyRight {
				g.player.ToRight()

				g.terraFormer.WithSnake(g.player)

				g.buff.Reset()
				g.terraFormer.Form(&g.buff)
				g.terraFormater.Print(&g.buff)
			}
			if event.Key() == tcell.KeyLeft {
				g.player.ToLeft()

				g.terraFormer.WithSnake(g.player)

				g.buff.Reset()
				g.terraFormer.Form(&g.buff)
				g.terraFormater.Print(&g.buff)
			}
		}
	}

}
