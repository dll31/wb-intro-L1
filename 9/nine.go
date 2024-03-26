// Задание: Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

// https://go.dev/blog/pipelines

package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func squares(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for inV := range in {
			out <- inV * inV
		}

		close(out)
	}()

	return out
}

func main() {

	for in := range squares(gen(2, 5, 10, 3, 6, 8)) {
		fmt.Println(in)
	}
}
