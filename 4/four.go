// Задание 4: Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func chanWorker(seqN int, dChan chan int, exit chan bool) {
	for {
		number, ok := <-dChan
		if !ok {
			// Прекращаем работу, если канал закрылся
			fmt.Printf("Stopping worker %d\n", seqN)
			exit <- true
			return
		}

		fmt.Printf("Worker %d get value %d\n", seqN, number)

	}
}

func chansMain() {

	var N int
	fmt.Println("Type workers number:")
	fmt.Scan(&N)

	dataChan := make(chan int)
	cleanupDone := make(chan bool, N)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	for i := 0; i < N; i++ {
		go chanWorker(i, dataChan, cleanupDone)
	}

	number := 0
	exit := false
	for !exit {
		select {
		case <-quit:
			fmt.Println("Get quit signal")
			// Закрываем канал, чтобы уведомить воркеры о прекращении работы
			close(dataChan)
			close(quit)
			exit = true
		default:
			time.Sleep(100 * time.Microsecond)
			number++
			dataChan <- number
		}
	}

	// Собираем подтверждения о закрытии воркеров
	for i := 0; i < N; i++ {
		<-cleanupDone
	}

	close(cleanupDone)
	fmt.Println("Main exit")
}

func wgWorker(wg *sync.WaitGroup, seqN int, dChan chan int) {
	defer wg.Done()

	for {
		number, ok := <-dChan
		if !ok {
			// Прекращаем работу, если канал закрылся
			fmt.Printf("Stopping worker %d\n", seqN)
			return
		}

		fmt.Printf("Worker %d get value %d\n", seqN, number)

	}
}

func wgMain() {
	var N int
	fmt.Println("Type workers number:")
	fmt.Scan(&N)

	dataChan := make(chan int)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go wgWorker(&wg, i, dataChan)
	}

	number := 0
	exit := false
	for !exit {
		select {
		case <-quit:
			fmt.Println("Get quit signal")
			// Закрываем канал, чтобы уведомить воркеры о прекращении работы
			close(dataChan)
			close(quit)
			exit = true
		default:
			time.Sleep(100 * time.Microsecond)
			number++
			dataChan <- number
		}
	}
	wg.Wait()
	fmt.Println("Main exit")
}
