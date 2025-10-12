package main

import "fmt"

type user struct {
	email    string
	nickname string
}

func createUser(email string, nickname string) (*user, error) {
	return &user{email, nickname}, nil
}

func main() {
	// Task A
	user, error := createUser("Mark", "example@example.com")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(*user)
	}
	fmt.Println("----------")

	// Task B
	_, otherError := createUser("Mark", "example@example.com")
	if otherError != nil {
		fmt.Println(otherError)
	}
}
