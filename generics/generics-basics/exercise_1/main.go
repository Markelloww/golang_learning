package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func marshal[T any](val T) ([]byte, error) {
	values, err := json.Marshal(val)

	if err != nil {
		return nil, err
	}

	return values, err
}

func main() {
	usr := User{
		Name: "Mark",
		Age:  20,
	}

	values, err := marshal(usr)

	fmt.Println(string(values))
	if err != nil {
		fmt.Println(err)
	}
}
