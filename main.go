package main

import (
	"dynimic-data-structure/figures"
	"fmt"
)

func main() {
	list := &LinkedList{}
	tree := &BinaryTree{}

	square := figures.NewSquare(5)
	square.SetColor("Red")
	square.SetRotation(450)

	rectangle := figures.NewRectangle(4, 6)
	rectangle.SetColor("Blue")
	rectangle.SetRotation(90)

	circle := figures.NewCircle(3)
	circle.SetColor("Green")
	circle.SetRotation(-30)

	triangle := figures.NewTriangle(4, 2)
	triangle.SetColor("Yellow")
	triangle.SetRotation(180)

	rhombus := figures.NewRhombus(6, 4)
	rhombus.SetColor("Orange")
	rhombus.SetRotation(720)

	shapes := []figures.Shape{square, rectangle, circle, triangle, rhombus}

	list.AddFirst(shapes[3])
	list.AddLast(shapes[0])
	list.AddLast(shapes[1])
	list.AddLast(shapes[2])
	list.AddLast(shapes[4])

	for _, shape := range shapes {
		tree.Insert(shape)
	}

	fmt.Println("Список фигур:")
	list.Print()

	fmt.Println("\nПлощади фигур:")
	list.PrintAreas()

	fmt.Println("\nДвоичное дерево фигур по площади:")
	tree.PrintInOrder()
}
