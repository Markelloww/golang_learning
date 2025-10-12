package main

import "fmt"

func main() {
	// Task A
	number := 20
	fmt.Println("Address:", &number)
	fmt.Println("Value:", number)
	fmt.Println("----------")

	// Task B
	pointer := &number
	fmt.Println("Address of main:", &number)
	fmt.Println("Address of pointer:", &pointer)
	fmt.Println("----------")
	fmt.Println("Value of main:", number)
	fmt.Println("Value of pointer:", *pointer)
}
