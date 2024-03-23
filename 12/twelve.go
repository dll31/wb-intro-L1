// Задание: Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	set := make(map[string]struct{})

	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, item := range seq {
		if _, exist := set[item]; !exist {
			set[item] = struct{}{}
		}
	}

	for key := range set {
		fmt.Println(key)
	}
}
