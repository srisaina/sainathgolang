package main

import "fmt"

type address struct {
	Name    string
	city    string
	pincode string
}

func main() {
	var a address
	fmt.Println(a)
	a1 := address{"sainath", "texas", "frisco"}
	fmt.Println("address:", a1)

}
