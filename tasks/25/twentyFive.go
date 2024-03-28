// Задача: Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func MySleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Before sleep")
	MySleep(10 * time.Second)
	fmt.Println("After sleep")
}
