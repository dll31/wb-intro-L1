// Задание: Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a, b := 3, 8

	// меняемся значениями
	a, b = b, a
	fmt.Println("a: ", a, " b: ", b)

	c, d := new(int), new(int)
	*c, *d = 10, 20

	// меняемся указателями на значения
	c, d = d, c

	fmt.Println("*c: ", *c, " *d: ", *d)

}
