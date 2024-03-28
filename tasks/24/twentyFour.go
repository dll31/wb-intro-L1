// Задача: Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func New(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func distanceBetween(p1 *Point, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}

func main() {
	p1 := New(10, 20)
	p2 := New(3, 3)
	fmt.Println(distanceBetween(p1, p2))
}
