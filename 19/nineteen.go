// Задание: Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

// "Hello, 世界"

package main

import (
	"fmt"
)

func reverseString(str string) (result string) {
	for _, value := range str {
		result = string(value) + result
	}
	return result
}

func main() {
	fmt.Println(reverseString("Hello, 世界"))
}
