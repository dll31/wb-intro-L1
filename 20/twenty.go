// Задание: 20.	Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog  sun"

	splittedString := strings.Split(str, " ")

	var builder strings.Builder

	builder.Grow(len(str))

	for i := len(splittedString) - 1; i >= 0; i-- {
		if splittedString[i] == "" {
			continue
		}

		builder.WriteString(splittedString[i])
		if i != 0 {
			builder.WriteString(" ")
		}
	}

	fmt.Println(builder.String())

}
