package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {
	mark := user{"Mark", "example@example.com", 20}
	changeAge(&mark, 21)
	fmt.Println(mark)
}

func changeAge(person *user, age int) {
	person.age = age
}
