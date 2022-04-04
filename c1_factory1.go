package main

import "fmt"

func main() {
	var num1, num2, res int
	var op string
	fmt.Println("Please enter 2 numbers: ")
	fmt.Scanln(&num1, &num2)

	fmt.Println("Please enter an operator: ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	}

	fmt.Printf("res is %d\n", res)

	return
}
