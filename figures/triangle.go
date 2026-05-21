package figures

import (
	"fmt"
)

type Triangle struct {
	BaseShape
	base   float64
	height float64
}

func NewTriangle(base, height float64) *Triangle {
	return &Triangle{base: base, height: height}
}

func (t Triangle) GetName() string {
	return "Triangle"
}

func (t Triangle) Draw() {
	fmt.Printf("%s: base=%.2f, height=%.2f, color=%s, rotation=%.2f, area=%.2f\n",
		t.GetName(), t.base, t.height, t.color, t.rotation, t.Area())
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}
