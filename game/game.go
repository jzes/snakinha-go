package game

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/jzes/snakinha-go/land"
	"github.com/jzes/snakinha-go/world"
	"github.com/jzes/snakinha-go/world/fruit"
)

type Game struct {
	screen tcell.Screen
	Buff   bytes.Buffer
	world  *world.World
	eaten  bool
}

func New(world *world.World) (*Game, error) {

	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err = screen.Init(); err != nil {
		return nil, err
	}

	screen.SetStyle(tcell.StyleDefault)
	g := Game{
		screen: screen,
		Buff:   bytes.Buffer{},
		world:  world,
		eaten:  false,
	}
	g.renewGame()

	return &g, nil
}

func (g Game) renewGame() {
	g.Buff.Reset()
	land.Form(&g.Buff, g.world)
	land.Print(&g.Buff, g.screen)
}

func (g Game) Loop() {

	land.Form(&g.Buff, g.world)

	var event tcell.Event
	go func() {
		for {
			g.screen.Show()
			event = g.screen.PollEvent()
		}
	}()

	for {
		switch event := event.(type) {
		case *tcell.EventResize:
			g.screen.Sync()
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyDown:
				g.world.Player.ToDown()
			case tcell.KeyUp:
				g.world.Player.ToUp()
			case tcell.KeyLeft:
				g.world.Player.ToLeft()
			case tcell.KeyRight:
				g.world.Player.ToRight()
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}
		}
		if g.world.Player.Ate(*g.world.Fruit) {
			g.world.Player.Grow1()
			newFruit := fruit.NewApple(
				g.world.Board.SizeX,
				g.world.Board.SizeY,
			)
			g.world.Fruit = &newFruit
		}

		if g.world.Player.AteItSelf() ||
			g.world.Player.ItsTheWall(g.world.Board) {
			fmt.Println("You Died")
			return
		}
		g.renewGame()
		g.screen.Show()
		time.Sleep(500 * time.Millisecond)
	}
}
