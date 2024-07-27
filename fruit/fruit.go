package fruit

type Fruit struct {
	X int
	Y int
}

type Fruits []Fruit

func (fs Fruits) CheckFruits(x, y int) bool {
	for _, f := range fs {
		if f.X == x && f.Y == y {
			return true
		}
	}
	return false
}
