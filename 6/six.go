// Задача: Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	// остановка с помощью канала
	ch := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func(exit chan struct{}) {
		select {
		case <-exit:
			fmt.Println("exit from goroutine")
			wg.Done()
			return
		}
	}(ch)
	time.Sleep(time.Second)
	close(ch)
	wg.Wait()

	// Завершение при помощи cancel функции контекста
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("exit from goroutine")
			wg.Done()
		}
	}(ctx)

	time.Sleep(time.Second)
	cancel()
	wg.Wait()

	timeout := time.Duration(3) * time.Second
	wg.Add(1)
	go func(timeout time.Duration) {
		select {
		case <-time.After(timeout):
			fmt.Println("Timeout")
			wg.Done()
		}
	}(timeout)

	wg.Wait()

}
