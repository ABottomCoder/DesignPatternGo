package main

import (
	"fmt"
	"strconv"
)

func ReadGoods() (prices []int) {
	fmt.Println("Please enter the price of goods: ")
	for {
		var s string
		fmt.Scan(&s)
		if s == "E" {
			break
		} else {
			price, _ := strconv.Atoi(s)
			prices = append(prices, price)
		}
	}
	fmt.Printf("prices of goods is: %v\n", prices)

	return
}

func GetTotalPrice(prices []int) (payPrice int) {
	for _, price := range prices {
		payPrice += price
	}

	return
}

func main() {
	prices := ReadGoods()
	payPrice := GetTotalPrice(prices)
	fmt.Printf("You should pay %d yuan\n", payPrice)

	return
}
