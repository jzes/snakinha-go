package snake

type BodySection struct {
	X int
	Y int
}

type Body []BodySection

func (b *Body) WithSection(bs BodySection) {
	*b = append(*b, bs)
}
