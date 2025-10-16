package main

import (
	"fmt"
	"math/rand"
)

const (
	grCount = 100
)

func main() {
	values := make(chan int, grCount)

	for gr := 0; gr < grCount; gr++ {
		go func() {
			values <- rand.Intn(1000)
		}()
	}

	wait := grCount

	var nums []int
	for wait > 0 {
		nums = append(nums, <-values)
		wait--
	}

	fmt.Println(nums)
}
