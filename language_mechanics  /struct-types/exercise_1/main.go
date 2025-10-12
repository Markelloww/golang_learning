package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {
	person := user{"Mark", "example@example.com", 20}
	fmt.Println(person)
}
