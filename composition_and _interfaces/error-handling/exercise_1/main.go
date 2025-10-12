package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidValue   = errors.New("Invalid Value")
	ErrAmountTooLarge = errors.New("Amount To Large")
)

func checkAmount(f float64) error {
	switch {
	case f == 0:
		return ErrInvalidValue
	case f > 1000:
		return ErrAmountTooLarge
	}
	return nil
}

func main() {
	if err := checkAmount(0); err != nil {
		switch err {
		case ErrInvalidValue:
			fmt.Println("Invalid value")
			return
		case ErrAmountTooLarge:
			fmt.Println("Value is too large")
			return
		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Everything is good")
}
