package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	mutex   sync.Mutex
	numbers []int
)

func main() {
	const grs = 3

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			random(10)
			wg.Done()
		}()
	}

	wg.Wait()

	for i, number := range numbers {
		fmt.Println(i, number)
	}
}

func random(amount int) {
	for i := 0; i < amount; i++ {
		n := rand.Intn(100)
		mutex.Lock()
		{
			numbers = append(numbers, n)
		}
		mutex.Unlock()
	}
}
