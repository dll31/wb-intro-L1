// Задание: Удалить i-ый элемент из слайса.

package main

import (
	"fmt"
	"slices"
)

func main() {
	sl := []int{20, 11, 323, 1, 0, 9, 34}
	i := 3

	fmt.Println(slices.Delete(sl, i, i+1))

	/* func Delete[S ~[]E, E any](s S, i, j int) S {
		_ = s[i:j:len(s)] // bounds check

		if i == j {
			return s
		}

		oldlen := len(s)
		s = append(s[:i], s[j:]...)
		clearSlice(s[len(s):oldlen]) // zero/nil out the obsolete elements, for GC
		return s
	} */
}
