// Задание: Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

/*
go func(dataChan chan int) {
	for {
		select {
		case time.After(time.Second * time.Duration(N)): // ждем дедлайна
			fmt.Println("Reciever: Deadline!")
			return
		case v, ok := <-dataChan: // или ждем данных в канал
			if ok {
				fmt.Println("Get some value from channel: ", v)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}
}(dataChan)
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	N := 0
	fmt.Printf("Enter number of seconds: ")
	fmt.Scan(&N)

	dataChan := make(chan int)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(N)))
	defer cancel()

	go func(ctx context.Context, dataChan chan int) {
		for {
			select {
			case <-ctx.Done(): // ждем дедлайна
				fmt.Println("Reciever: Deadline!")
				return
			case v, ok := <-dataChan: // или ждем данных в канал
				if ok {
					fmt.Println("Get some value from channel: ", v)
				} else {
					fmt.Println("Channel closed")
					return
				}
			}
		}
	}(ctx, dataChan)

	i := 0
	exit := false
	for !exit {
		i++
		select {
		case <-ctx.Done():
			fmt.Println("Sender: Deadline!")
			close(dataChan)
			cancel()
			exit = true
		default:
			dataChan <- i
		}
	}

	fmt.Println("Exit")

}
