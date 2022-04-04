package main

import "fmt"

type SchoolGirl struct {
	name string
}

type GiveGift interface {
	GiveDoll()
	GiveFlowers()
}

type Pursuit struct {
	mm SchoolGirl
}

func (p Pursuit) GiveDoll() {
	fmt.Printf("%s, give you doll\n", p.mm.name)
}

func (p Pursuit) GiveFlowers() {
	fmt.Printf("%s, give you flowers\n", p.mm.name)
}

type Proxy struct {
	gg Pursuit
}

func (p Proxy) GiveDoll() {
	p.gg.GiveDoll()
}

func (p Proxy) GiveFlowers() {
	p.gg.GiveFlowers()
}

func main() {
	jiaojiao := SchoolGirl{name: "Jiao"}
	zhuojiayi := Pursuit{mm: jiaojiao}
	daili := Proxy{gg: zhuojiayi}

	daili.GiveDoll()
	daili.GiveFlowers()
}
