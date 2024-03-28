// Задание: Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

package main

import "fmt"

func main() {
	var nums = [...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	var spreadValues = make(map[int][]float32)

	for _, val := range nums {
		var quotient = int(val) / 10 * 10
		spreadValues[quotient] = append(spreadValues[quotient], val)
	}

	fmt.Printf("%+v", spreadValues)
}
