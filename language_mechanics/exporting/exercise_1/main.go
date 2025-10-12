package main

import (
	"awesomeProject/language_mechanics/exporting/exercise_1/toy"
	"fmt"
)

func main() {
	var car = toy.New("Car", 7)
	car.UpdateOnHand(2)
	car.UpdateSold(14)
	fmt.Println("Name:", car.Name)
	fmt.Println("Weight:", car.Weight)
	fmt.Println("OnHand:", car.OnHand())
	fmt.Println("Sold:", car.Sold())
}
