package main

import "fmt"

type SchoolGirl struct {
	name string
}

type Pursuit struct {
	mm SchoolGirl
}

func (p Pursuit) GiveDolls() {
	fmt.Printf("%s, give you doll\n", p.mm.name)
}

func (p Pursuit) GiveFlowers() {
	fmt.Printf("%s, give you flowers\n", p.mm.name)
}

func main() {
	jiaojiao := SchoolGirl{name: "Jiao"}
	zhuojiayi := Pursuit{mm: jiaojiao}
	zhuojiayi.GiveDolls()
	zhuojiayi.GiveFlowers()
}
