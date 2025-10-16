package main

import "fmt"

type copyer interface {
	string | int
}

func copyfy[T copyer](val []T) []T {
	tmp := make([]T, len(val))
	copy(tmp, val)
	return tmp
}

func main() {
	slc := []string{"val1", "val2", "val3"}

	slc2 := copyfy(slc)

	fmt.Println(slc)
	fmt.Println(slc2)

	slc3 := []int{1, 2, 3}
	slc4 := copyfy(slc3)

	fmt.Println(slc3)
	fmt.Println(slc4)
}
