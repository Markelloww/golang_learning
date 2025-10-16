package main

import (
	"fmt"
	"sync"
)

func main() {
	gr := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		goroutine("Bill", gr)
		wg.Done()
	}()

	go func() {
		goroutine("Joan", gr)
		wg.Done()
	}()

	gr <- 1

	wg.Wait()
}

func goroutine(name string, gr chan int) {
	for {
		value, ok := <-gr
		if !ok {
			fmt.Printf("Горутина: %s отработала\n", name)
			return
		}

		fmt.Printf("Горутина: %s, итерация: %d\n", name, value)

		if value == 10 {
			close(gr)
			fmt.Printf("Горутина: %s,  закончила свое выполнение на 10 итерации\n", name)
			return
		}

		gr <- value + 1
	}
}
