package figures

type Shape interface {
	Draw()
	GetName() string
	Area() float64
}
