package main

import (
	"log"

	"github.com/jzes/snakinha-go/game"
	"github.com/jzes/snakinha-go/world"
	"github.com/jzes/snakinha-go/world/board"
	"github.com/jzes/snakinha-go/world/fruit"
	"github.com/jzes/snakinha-go/world/snake"
)

func main() {

	body := snake.Body{}
	body.WithSection(snake.BodySection{X: 9, Y: 1})
	body.WithSection(snake.BodySection{X: 10, Y: 1})
	body.WithSection(snake.BodySection{X: 11, Y: 1})
	body.WithSection(snake.BodySection{X: 12, Y: 1})
	player1 := snake.New(body, "o")

	board1 := board.New(" ", 40, 20)

	fruit1 := fruit.NewApple(40, 20)

	world := world.World{
		Board:  &board1,
		Player: &player1,
		Fruit:  &fruit1,
	}

	g, err := game.New(&world)
	if err != nil {
		log.Fatal("erro ao iniciar o game")
	}
	g.Loop()

}
