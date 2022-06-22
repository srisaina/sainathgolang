package main

import (
	"fmt"
)

func main() {
	sampleslice1 := []int{}
	fmt.Println(len(sampleslice1))
	fmt.Println(cap(sampleslice1))
	fmt.Println(sampleslice1)

	sampleslice2 := []string{"sa", "ge", "it", "golang"}
	fmt.Println(len(sampleslice2))
	fmt.Println(cap(sampleslice2))
	fmt.Println(sampleslice2)
}
