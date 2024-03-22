// Задание 8: Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func main() {
	var number int64 = -35 //rand.Int63()
	// if rand.Intn(2) == 1 {
	// 	number *= -1
	// }

	var bitNumber, bitValue int
	fmt.Printf("Enter bit number [1; 64]: ")
	fmt.Scan(&bitNumber)
	fmt.Printf("Enter bit value (0 or 1): ")
	fmt.Scan(&bitValue)

	fmt.Printf("%064b\n", number)

	if number >= 0 {
		if bitValue == 0 {
			number &^= (1 << (bitNumber - 1))
		} else if bitValue == 1 {
			number |= (1 << (bitNumber - 1))
		}
	} else {
		if bitValue == 0 {
			number |= (1 << (bitNumber - 1))
		} else if bitValue == 1 {
			number &^= (1 << (bitNumber - 1))
		}
	}
	fmt.Printf("%064b\n", number)
}
