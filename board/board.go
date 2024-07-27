package board

type Board struct {
	SizeX int
	SizeY int
}

type OutOfBoardBounds string

func (e *OutOfBoardBounds) Error() string {
	*e = "Error, some iten is out of board boundaries"
	return string(*e)
}
