package main

import (
	"fmt"
	"sainathgolang/car/app"
)

func main() {
	//	a := app.Car{Engine: "v6", Tires: 3, SteeringWheels: "fabric", Paint: "blue"}
	mycar := app.FordCar("v8", 4, "leather", "white")
	fmt.Println(mycar)

}
