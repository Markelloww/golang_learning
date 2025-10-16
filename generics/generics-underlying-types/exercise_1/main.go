package main

import "fmt"

type keymap[T any] map[string]T

func (km keymap[T]) set(key string, val T) {
	km[key] = val
}

func (km keymap[T]) get(key string) (T, bool) {
	value, flag := km[key]
	if !flag {
		var empty T
		return empty, false
	}
	return value, true
}

func main() {
	km := make(keymap[int])

	km.set("jack", 10)
	km.set("jill", 20)

	jack, found := km.get("jack")
	if !found {
		fmt.Println("Jack не был найден")
	} else {
		fmt.Println("Jack был найден и хранит значение:", jack)
	}

	jill, found := km.get("jill")
	if !found {
		fmt.Println("Jill не была найдена")
	} else {
		fmt.Println("Jill была найдена и хранит значение:", jill)
	}
}
