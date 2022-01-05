package main

import "fmt"

type PlaneFigure interface {
	countWidth() int
}

type Square struct {
	Length int
}

func (square Square) countWidth() int {
	return square.Length * square.Length
}

type Rectangle struct {
	Length int
	Width  int
}

func (rectangle Rectangle) countWidth() int {
	return rectangle.Length * rectangle.Width
}

func main() {
	square1 := Square{10}
	rectangle1 := Rectangle{10, 5}

	square1_width := countAnyWidth(square1)
	rectangle1_width := countAnyWidth(rectangle1)

	fmt.Printf("Width of square 1: %d\nWidth of rectangle 1: %d\n", square1_width, rectangle1_width)
}

func countAnyWidth(figure PlaneFigure) int {
	return figure.countWidth()
}
