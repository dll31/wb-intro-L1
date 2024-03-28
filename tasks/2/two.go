// Задание 2: Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var numbers = []int{2, 4, 6, 8, 10}

	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			fmt.Println(math.Pow(float64(n), 2))
		}(num)
	}
	wg.Wait()

}
