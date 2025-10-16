package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for count := 100; count >= 0; count-- {
			fmt.Printf("A: %d\n", count)
		}
		wg.Done()
	}()

	go func() {
		for count := 0; count <= 100; count++ {
			fmt.Printf("B: %d\n", count)
		}

		wg.Done()
	}()

	fmt.Println("Ждем потоки")
	wg.Wait()

	fmt.Println("Все потоки выполнили свою работу")
}
