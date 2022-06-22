package app

type Car struct {
	Engine         string
	Tires          int
	SteeringWheels string
	Paint          string
}

func FordCar(fcar Car) { //i am building a ford car with car design with the datatype struct and name of the varaible is car
	fcar.Engine = "v8"
	fcar.Tires = 4
	fcar.SteeringWheels = "leather"
	fcar.Paint = "red"

}
