package figures

import "math"

type BaseShape struct {
	color    string
	rotation float64
}

func (b *BaseShape) GetColor() string      { return b.color }
func (b *BaseShape) SetColor(color string) { b.color = color }

func (b *BaseShape) GetRotation() float64 {
	return b.rotation
}

func (b *BaseShape) SetRotation(rotation float64) {
	rotation = math.Mod(rotation, 360.0)
	if rotation < 0 {
		rotation += 360
	}
	b.rotation = rotation
}
