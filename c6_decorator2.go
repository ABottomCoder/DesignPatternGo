package main

import "fmt"

type Person struct{
	name string
}

func (p Person)Show(){
	fmt.Printf("decorate person: %s", p.name)
}

type Finery interface{
	Show()
}

type Tshirts struct{
}
func (t Tshirts)Show(){
	fmt.Println("wear T-shirts!!!")
}

type BigTrouser struct{
}

func (b BigTrouser)Show(){
	fmt.Println("wear big trouser!!!")
}

func main(){
	xc := Person{name : "Cai"}

	fmt.Println("first collection: ")
	dtx := Tshirts{}
	kk := BigTrouser{}

	dtx.Show()
	kk.Show()
	xc.Show()
}
