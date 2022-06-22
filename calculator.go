package main

import "fmt"

func main() {

	var a, b float64
	var oper string

	fmt.Print("Enter the first number : ")
	fmt.Scanln(&a)

	fmt.Print("Enter the second number : ")
	fmt.Scanln(&b)

	fmt.Print("Enter the operator (+ - * /) :")
	fmt.Scan(&oper)

	switch oper {

	case "+":
		fmt.Printf("%f %s %f = %f", a, oper, b, a+b)

	case "-":
		fmt.Printf("%f %s %f = %f", a, oper, b, a-b)

	case "*":
		fmt.Printf("%f %s %f = %f", a, oper, b, a*b)

	case "/":
		if b == 0 {
			fmt.Println("Divided by Zero")
		} else {
			fmt.Printf("%f %s %f = %f", a, oper, b, a/b)
		}
	}
}
