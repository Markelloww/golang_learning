package main

import (
	"errors"
	"fmt"
)

type appError struct {
	err     error
	message string
	code    int
}

func (app *appError) Error() string {
	return fmt.Sprintf("Error[%s]; Message[%s]; Code[%d]", app.err, app.message, app.code)
}

func (app *appError) Temporary() bool {
	return app.code != 9
}

type temporary interface {
	Temporary() bool
}

func main() {
	if err := checkFlag(false); err != nil {
		switch e := err.(type) {
		case temporary:
			fmt.Println(err)
			if !e.Temporary() {
				fmt.Println("Critical Error")
			}
		default:
			fmt.Println(err)
		}
	}
}

func checkFlag(t bool) error {
	if !t {
		return &appError{errors.New("flag false"), "Flag was false", 9}
	}

	return errors.New("flag true")
}
