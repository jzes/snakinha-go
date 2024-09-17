package board

type Board struct {
	Sprite string
	SizeX  int
	SizeY  int
}

func New(s string, x, y int) Board {
	return Board{
		SizeY:  y,
		SizeX:  x,
		Sprite: s,
	}
}

type OutOfBoardBounds string

func (e *OutOfBoardBounds) Error() string {
	*e = "Error, some iten is out of board boundaries"
	return string(*e)
}
