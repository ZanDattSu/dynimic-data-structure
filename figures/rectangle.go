package figures

import (
	"fmt"
)

type Rectangle struct {
	BaseShape
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r Rectangle) GetName() string {
	return "Rectangle"
}

func (r Rectangle) Draw() {
	fmt.Printf("%s: width=%.2f, height=%.2f, color=%s, rotation=%.2f, area=%.2f\n",
		r.GetName(), r.width, r.height, r.color, r.rotation, r.Area())
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
