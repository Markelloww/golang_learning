package main

import "fmt"

func main() {
	var empty [5]string
	names := [5]string{"Mark", "Alina", "Vlad", "Alex", "Ann"}

	empty = names

	for i, name := range empty {
		fmt.Println(name, &empty[i])
	}

	fmt.Println("----------")

	for i, name := range names {
		fmt.Println(name, &names[i])
	}
}
