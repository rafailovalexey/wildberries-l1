package main

import "fmt"

/*
	№ 21 (1-ое решение)

	Реализовать паттерн «адаптер» на любом примере.
*/

type FiguresInterface interface {
	GetPerimeter() int
	GetArea() int
}

type Square struct {
	Side int
}

var _ FiguresInterface = (*Square)(nil)

func (s *Square) GetPerimeter() int {
	return 4 * s.Side
}

func (s *Square) GetArea() int {
	return s.Side * s.Side
}

type SquareAdapter struct {
	Square
}

func NewSquareAdapter(side int) *SquareAdapter {
	return &SquareAdapter{Square{Side: side}}
}

type Rectangle struct {
	Width  int
	Height int
}

var _ FiguresInterface = (*Rectangle)(nil)

func (r *Rectangle) GetPerimeter() int {
	return 2 * (r.Width + r.Height)
}

func (r *Rectangle) GetArea() int {
	return r.Width * r.Height
}

type RectangleAdapter struct {
	Rectangle
}

func NewRectangleAdapter(width int, height int) *RectangleAdapter {
	return &RectangleAdapter{Rectangle{Width: width, Height: height}}
}

func main() {
	square := NewSquareAdapter(5)
	fmt.Printf("square perimeter %d\n", square.GetPerimeter())
	fmt.Printf("square area %d\n", square.GetArea())

	rectangle := NewRectangleAdapter(10, 5)
	fmt.Printf("rectangle perimeter %d\n", rectangle.GetPerimeter())
	fmt.Printf("rectangle area %d\n", rectangle.GetArea())
}
