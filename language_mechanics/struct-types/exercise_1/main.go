package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {
	// Task A
	mark := user{"Mark", "example@example.com", 20}
	fmt.Println(mark)

	// Task B
	alina := struct {
		name  string
		email string
		age   int
	}{
		name:  "Aline",
		email: "example@example.com",
		age:   20,
	}
	fmt.Println(alina)
}
