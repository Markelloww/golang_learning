package main

import (
	"fmt"
)

func main() {
	// Task A
	var values []int

	for i := 0; i < 10; i++ {
		values = append(values, i)
	}

	for _, value := range values {
		fmt.Print(value, " ")
	}
	fmt.Println("\n----------")

	// Task B
	names := []string{"Mark", "Alina", "Vlad", "Alex", "Ann"}
	fmt.Println(names)

	slice := names[1:3]
	for i, el := range slice {
		fmt.Printf("Index: %d, value: %s\n", i, el)
	}
}
