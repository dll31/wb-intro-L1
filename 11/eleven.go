// Задание: Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersection[T []any](f, s T) T {
	// Создаем map[значение]количество повторений
	i := make(map[any]int)

	for _, value := range f {
		i[value]++
	}

	for _, value := range s {
		i[value]++
	}

	// Для всех повторяющихся значений количество повторений будет 2 или больше
	out := make([]any, 0)
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

	// []int does not satisfy []any ([]int missing in []any)
	// fmt.Printf("%+v\n", intersection(first, second))
	/*  Can I convert a []T to an []interface{}?
	Not directly. It is disallowed by the language specification because the two types do not have the same representation in memory. It is necessary to copy the elements individually to the destination slice.
	*/

	tFirst := make([]any, len(first))
	for i, v := range first {
		tFirst[i] = v
	}

	tSecond := make([]any, len(second))
	for i, v := range second {
		tSecond[i] = v
	}

	fmt.Printf("%+v\n", intersection(tFirst, tSecond))
}
