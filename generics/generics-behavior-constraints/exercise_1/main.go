package main

import (
	"encoding/json"
	"fmt"
)

func marshal[T json.Marshaler](val T) ([]byte, error) {
	data, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return data, err
}

type user struct {
	name  string
	email string
}

func (usr user) MarshalJSON() ([]byte, error) {
	val := fmt.Sprintf("{\"name\": %q, \"email\": %q}", usr.name, usr.email)
	return []byte(val), nil
}

func main() {
	usr := user{
		name:  "Mark",
		email: "example@example.com",
	}

	val, err := marshal(usr)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(val))
	}
}
