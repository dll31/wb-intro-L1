// Задание: Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

type counter struct {
	sync.RWMutex
	c int
}

func main() {
	var count = counter{
		c: 0,
	}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			incs := rand.Intn(15)
			fmt.Printf("Worker %d start. incs = %d\n", n, incs)
			for i := 0; i < incs; i++ {
				// increment
				count.Lock()
				count.c++
				count.Unlock()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("Final mutex count", count.c)

	var c int32 = 0
	var wg2 sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func(n int) {
			incs := rand.Intn(15)
			fmt.Printf("Worker %d start. incs = %d\n", n, incs)
			for i := 0; i < incs; i++ {
				// increment
				atomic.AddInt32(&c, 1)
			}
			wg2.Done()
		}(i)
	}

	wg2.Wait()
	fmt.Println("Final atomic count", c)
}
