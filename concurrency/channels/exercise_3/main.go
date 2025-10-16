package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	grCount = 100
)

func main() {
	values := make(chan int)

	var wg sync.WaitGroup
	wg.Add(grCount)

	for gr := 0; gr < grCount; gr++ {
		go func() {
			defer wg.Done()

			n := rand.Intn(1000)

			if n%2 == 0 {
				return
			}

			values <- n
		}()
	}

	go func() {
		wg.Wait()
		close(values)
	}()

	var nums []int
	for n := range values {
		nums = append(nums, n)
	}

	fmt.Printf("Результирующее кол-во: %d\n", len(nums))
	fmt.Println(nums)
}
