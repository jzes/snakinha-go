package fruit

type Gardem struct {
	Fruits Fruits
	Sprite string
}

func NewGardem(s string) Gardem {
	return Gardem{
		Sprite: s,
	}
}

func (g *Gardem) WithFruit(f Fruit) {
	g.Fruits = append(g.Fruits, f)
}
