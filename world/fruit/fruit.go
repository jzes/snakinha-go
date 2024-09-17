package fruit

import "math/rand/v2"

type Fruit struct { //a sprite fica aqui ? ou melhor colocar no gardem e mante-lo?
	Sprite string
	X      int
	Y      int
}

func NewApple(x, y int) Fruit {
	return Fruit{"*", rand.IntN(x), rand.IntN(y)}
}

func (f Fruit) Equals(x, y int) bool {
	return f.X == x && f.Y == y
}
