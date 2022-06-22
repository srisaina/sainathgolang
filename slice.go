package main

import (
	"fmt"
)

func main() {
	a := [6]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	myslice := a[1:3]
	fmt.Println("myslice =", myslice)
}
