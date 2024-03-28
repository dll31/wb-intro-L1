// Задание 3:Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var numbers = []int{2, 4, 6, 8, 10}
	outCh := make(chan int)

	// Выносим в отдельную горутину, чтобы не упасть в deadlock
	go func() {
		for _, num := range numbers {
			wg.Add(1)

			go func(n int) {
				defer wg.Done()
				outCh <- int(math.Pow(float64(n), 2))
			}(num)
		}
		wg.Wait()
		close(outCh)
	}()

	var res int = 0
	for item := range outCh {
		res += item
	}

	fmt.Println(res)

}
