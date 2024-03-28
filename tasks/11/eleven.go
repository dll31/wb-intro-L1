// Задание: Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersection[T comparable](f, s []T) []T {
	// Создаем map[значение]количество повторений
	i := make(map[T]int)

	for _, value := range f {
		i[value]++
	}

	for _, value := range s {
		i[value]++
	}

	// Для всех повторяющихся значений количество повторений будет 2 или больше
	out := make([]T, 0)
	for k, v := range i {
		if v >= 2 {
			out = append(out, k)
		}
	}
	return out
}

func main() {
	first := []int{1, 3, 5, 7, 20, 33, 2, 12}
	second := []int{9, 11, 7, 33, 48, 13, 2}

	fmt.Printf("%+v\n", intersection(first, second))
}
