package figures

import (
	"fmt"
)

type Square struct {
	BaseShape
	side float64
}

func NewSquare(side float64) *Square {
	return &Square{side: side}
}

func (s Square) GetName() string {
	return "Square"
}

func (s Square) Draw() {
	fmt.Printf("%s: side=%.2f, color=%s, rotation=%.2f, area=%.2f\n",
		s.GetName(), s.side, s.color, s.rotation, s.Area())
}

func (s Square) Area() float64 {
	return s.side * s.side
}
