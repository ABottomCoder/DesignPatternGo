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

type Fund struct {
	s1 Stock1
	s2 Stock2
	d1 NationalDebt1
}

func (f Fund) BuyFund() {
	f.s1.Buy()
	f.s2.Buy()
	f.d1.Buy()
}

func (f Fund) SellFund() {
	f.s1.Sell()
	f.s2.Sell()
	f.d1.Sell()
}

func main() {
	f := Fund{}
	f.BuyFund()
	f.SellFund()
}
