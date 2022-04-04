package main

import "fmt"

type SchoolGirl struct {
	name string
}

type Proxy struct {
	mm SchoolGirl
}

func (p Proxy) GiveDolls() {
	fmt.Printf("%s, give you doll\n", p.mm.name)
}

func (p Proxy) GiveFlowers() {
	fmt.Printf("%s, give you flowers\n", p.mm.name)
}

func main() {
	jiaojiao := SchoolGirl{name: "Jiao"}
	daili := Proxy{mm: jiaojiao}
	daili.GiveDolls()
	daili.GiveFlowers()
}
