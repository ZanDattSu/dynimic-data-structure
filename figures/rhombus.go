package figures

import (
	"fmt"
)

type Rhombus struct {
	BaseShape
	diagonal1 float64
	diagonal2 float64
}

func NewRhombus(diagonal1, diagonal2 float64) *Rhombus {
	return &Rhombus{diagonal1: diagonal1, diagonal2: diagonal2}
}

func (r Rhombus) GetName() string {
	return "Rhombus"
}

func (r Rhombus) Draw() {
	fmt.Printf("%s: diagonal1=%.2f, diagonal2=%.2f, color=%s, rotation=%.2f, area=%.2f\n",
		r.GetName(), r.diagonal1, r.diagonal2, r.color, r.rotation, r.Area())
}

func (r Rhombus) Area() float64 {
	return r.diagonal1 * r.diagonal2 / 2
}
