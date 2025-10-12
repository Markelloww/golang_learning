package main

import "fmt"

func main() {
	people := make(map[string]int)

	people["Mark"] = 20
	people["Alina"] = 20
	people["Alex"] = 22
	people["Anna"] = 19
	people["Max"] = 21

	for key, value := range people {
		fmt.Printf("Name: %s, age: %d\n", key, value)
	}
}
