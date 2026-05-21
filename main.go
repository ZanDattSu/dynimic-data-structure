package main

import (
	"dynimic-data-structure/figures"
	"fmt"
	"strings"
)

func printHeader(title string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60))
}

func main() {
	// =========================================================
	// 1. СОЗДАНИЕ ГЕОМЕТРИЧЕСКИХ ФИГУР
	// =========================================================
	printHeader("1. Создание геометрических фигур")

	square := figures.NewSquare(5)
	square.SetColor("Red")
	square.SetRotation(450) // нормализуется в 90°

	rectangle := figures.NewRectangle(4, 6)
	rectangle.SetColor("Blue")
	rectangle.SetRotation(90)

	circle := figures.NewCircle(3)
	circle.SetColor("Green")
	circle.SetRotation(-30) // нормализуется в 330°

	triangle := figures.NewTriangle(4, 2)
	triangle.SetColor("Yellow")
	triangle.SetRotation(180)

	rhombus := figures.NewRhombus(6, 4)
	rhombus.SetColor("Orange")
	rhombus.SetRotation(720) // нормализуется в 0°

	shapes := []figures.Shape{square, rectangle, circle, triangle, rhombus}

	for _, s := range shapes {
		s.Draw()
	}

	// =========================================================
	// 2. ДЕМОНСТРАЦИЯ РАБОТЫ ДВУСВЯЗНОГО СПИСКА
	// =========================================================
	printHeader("2. Двусвязный список (LinkedList)")

	list := &LinkedList{}

	triangleNode := list.AddFirst(triangle)
	list.AddLast(square)
	list.AddLast(rectangle)
	list.AddLast(circle)
	list.AddLast(rhombus)

	fmt.Println("Список после добавления:")
	list.Print()

	fmt.Println("\nПлощади фигур в списке:")
	list.PrintAreas()

	fmt.Println("\nУдаление треугольника из списка...")
	list.Remove(triangleNode)

	fmt.Println("\nСписок после удаления:")
	list.Print()

	// =========================================================
	// 3. ДЕМОНСТРАЦИЯ РАБОТЫ БИНАРНОГО ДЕРЕВА ПОИСКА
	// =========================================================
	printHeader("3. Бинарное дерево поиска (BinaryTree)")

	tree := &BinaryTree{}

	tree.Insert(rectangle)
	tree.Insert(rhombus)
	tree.Insert(square)
	tree.Insert(triangle)
	tree.Insert(circle)

	fmt.Println("In-order обход (фигуры упорядочены по площади):")
	tree.PrintInOrder()

	fmt.Printf("\nВысота дерева:   %d\n", tree.Height())
	fmt.Printf("Размер дерева:   %d\n", tree.Size())

	minShape := tree.FindMin()
	maxShape := tree.FindMax()
	if minShape != nil {
		fmt.Printf("Минимальная площадь: %s (S = %.2f)\n", minShape.GetName(), minShape.Area())
	}
	if maxShape != nil {
		fmt.Printf("Максимальная площадь: %s (S = %.2f)\n", maxShape.GetName(), maxShape.Area())
	}

	fmt.Println("\nПоиск фигур по площади:")
	searchAreas := []float64{25.0, 12.0, 100.0}
	for _, area := range searchAreas {
		found := tree.Search(area)
		if found != nil {
			fmt.Printf("  S = %.2f -> найдена: %s\n", area, found.GetName())
		} else {
			fmt.Printf("  S = %.2f -> не найдена\n", area)
		}
	}

	fmt.Println("\nПоиск фигур с площадью в диапазоне [10, 26]:")
	rangeShapes := tree.FindInRange(10, 26)
	for _, s := range rangeShapes {
		fmt.Printf("  %s: S = %.2f\n", s.GetName(), s.Area())
	}

	fmt.Println("\nПреобразование дерева в слайс (ToSlice):")
	slice := tree.ToSlice()
	for i, s := range slice {
		fmt.Printf("  [%d] %s, S = %.2f\n", i, s.GetName(), s.Area())
	}

	// =========================================================
	// 4. ЗАВЕРШЕНИЕ
	// =========================================================
	printHeader("Демонстрация завершена")
}
