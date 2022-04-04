package main

import "fmt"

type Operation interface {
	GetResult(num1 int, num2 int) int
}

type OperationAdd struct {
}

func (a OperationAdd) GetResult(num1 int, num2 int) int {
	return num1 + num2
}

type OperationSub struct {
}

func (a OperationSub) GetResult(num1 int, num2 int) int {
	return num1 - num2
}

type CalculatorFactory struct {
}

func (c CalculatorFactory) GetFactory(s string) Operation {
	switch s {
	case "+":
		return OperationAdd{}
	case "-":
		return OperationSub{}
	default:
		return nil
	}
}

func main() {
	var op string
	var num1, num2, res int
	fmt.Println("Please enter 2 numbers: ")
	fmt.Scanln(&num1, &num2)
	fmt.Println("Please enter an operator: ")
	fmt.Scanln(&op)

	factory := CalculatorFactory{}
	myFactory := factory.GetFactory(op)

	res = myFactory.GetResult(num1, num2)
	fmt.Printf("res is %d\n", res)

	return
}
