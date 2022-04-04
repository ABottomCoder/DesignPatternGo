package main

import "fmt"

type Stock1 struct {
}

func (s Stock1) Buy() {
	fmt.Println("Buy stock 1")
}

func (s Stock1) Sell() {
	fmt.Println("Sell stock 1")
}

type Stock2 struct {
}

func (s Stock2) Buy() {
	fmt.Println("Buy stock 2")
}

func (s Stock2) Sell() {
	fmt.Println("Sell stock 2")
}

type NationalDebt1 struct {
}

func (s NationalDebt1) Buy() {
	fmt.Println("Buy NationalDebt 1")
}

func (s NationalDebt1) Sell() {
	fmt.Println("Sell NationalDebt 1")
}

func main() {
	stock1 := Stock1{}
	stock2 := Stock2{}
	nationalDebt := NationalDebt1{}

	stock1.Buy()
	stock1.Sell()
	stock2.Buy()
	stock2.Sell()
	nationalDebt.Buy()
	nationalDebt.Sell()
}
