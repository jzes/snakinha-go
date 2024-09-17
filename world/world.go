package world

import (
	"github.com/jzes/snakinha-go/world/board"
	"github.com/jzes/snakinha-go/world/fruit"
	"github.com/jzes/snakinha-go/world/snake"
)

type World struct {
	Board  *board.Board
	Player *snake.Snake
	Fruit  *fruit.Fruit
}
