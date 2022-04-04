package main

import (
	"fmt"
	"math"
	"strconv"
)

type CashSuper interface {
	AcceptCash(total float64) float64
}

type CashNormal struct {
}

func (c CashNormal) AcceptCash(total float64) float64 {
	return total
}

type CashRebate struct {
	moneyRebate float64
}

func (c CashRebate) AcceptCash(total float64) float64 {
	return total * c.moneyRebate
}

type CashReturn struct {
	moneyCondition float64
	moneyReturn    float64
}

func (c CashReturn) AcceptCash(total float64) float64 {
	return total - c.moneyReturn*math.Floor(total/c.moneyCondition)
}

type CashContext struct {
	cs CashSuper
}

func (c *CashContext) SetGetSuper(method string) {
	switch method {
	case "normal":
		c.cs = CashNormal{}
	case "rebate":
		println("Please enter a rebate number: ")
		var rebate float64
		fmt.Scan(&rebate)
		c.cs = CashRebate{moneyRebate: rebate}
	case "return":
		println("Please enter condition and myReturn number: ")
		var condition, myReturn float64
		fmt.Scanln(&condition, &myReturn)
		c.cs = CashReturn{moneyCondition: condition, moneyReturn: myReturn}
	default:
		fmt.Println("unknown method")
	}
}

func (c CashContext) GetResult(money float64) float64 {
	return c.cs.AcceptCash(money)
}

func ReadGoodsFloat() (prices []float64) {
	fmt.Println("Please enter the price of goods: ")
	for {
		var s string
		fmt.Scan(&s)
		if s == "E" {
			break
		} else {
			price, _ := strconv.ParseFloat(s, 64)
			prices = append(prices, price)
		}
	}
	fmt.Printf("prices of goods is: %v\n", prices)

	return
}

func GetTotalPriceFloat(prices []float64) (payPrice float64) {
	for _, price := range prices {
		payPrice += price
	}

	return
}

func main() {
	prices := ReadGoodsFloat()
	total := GetTotalPriceFloat(prices)
	fmt.Println("Please enter your charge method: ")
	var method string
	fmt.Scan(&method)

	ctx := &CashContext{}
	ctx.SetGetSuper(method)

	charge := ctx.GetResult(total)
	fmt.Printf("finally you should pay: %.2f", charge)
	return
}
