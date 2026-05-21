package figures

import (
	"fmt"
	"math"
)

type Circle struct {
	BaseShape
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c Circle) GetName() string {
	return "Circle"
}

func (c Circle) Draw() {
	fmt.Printf("%s: radius=%.2f, color=%s, rotation=%.2f, area=%.2f\n",
		c.GetName(), c.radius, c.color, c.rotation, c.Area())
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
