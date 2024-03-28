// Задание: Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type concMap struct {
	sync.RWMutex
	m map[int]int
}

func main() {
	cmap := concMap{
		m: make(map[int]int),
	}
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			for j := i * 10; j < (i*10 + 10); j++ {
				cmap.Lock()
				cmap.m[j] = i
				cmap.Unlock()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	for k, v := range cmap.m {
		fmt.Println("Key: ", k, " Value: ", v)
	}
}
