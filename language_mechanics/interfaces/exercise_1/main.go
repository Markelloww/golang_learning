package main

import (
	"fmt"
)

type speaker interface {
	speak()
}

type english struct{}

type chinese struct{}

func (english) speak() {
	fmt.Println("Hello world!")
}

func (chinese) speak() {
	fmt.Println("你好世界")
}

func sayHello(sp speaker) {
	sp.speak()
}

func main() {
	// Task A
	var en speaker = english{}
	en.speak()
	var ch speaker = chinese{}
	ch.speak()
	fmt.Println("----------")

	// Task B
	sayHello(english{})
	sayHello(chinese{})
}
