package main

import "fmt"

type Person struct {
	name string
}

func (p Person)WearTshirts(){
	fmt.Println("wear a T-shirts!!!")
}

func (p Person)WearSuit(){
	fmt.Println("wear a suit!!!")
}

func main(){
	xc := Person{name:"Cai"}

	fmt.Println("first collocation: ")
	xc.WearTshirts()
	xc.WearSuit()
}
