package main

import "fmt"

func getRes(num1, num2 int, op string) (res int) {
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

	return
}

func main() {
	var num1, num2, res int
	var op string
	fmt.Println("Please enter 2 numbers: ")
	fmt.Scanln(&num1, &num2)

	fmt.Println("Please enter an operator: ")
	fmt.Scanln(&op)

	res = getRes(num1, num2, op)
	fmt.Printf("res is %d\n", res)

	return
}
