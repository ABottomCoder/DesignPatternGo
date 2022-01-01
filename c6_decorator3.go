package main

import "fmt"

type Base interface{
	Show()
}

type Person struct{
	name string
}

func (p Person)Show(){
	fmt.Printf("decorated %s", p.name)
}

type Cloth interface{
	Base
	Decorate()
}

type Finery struct{
	base Base
}

func (f *Finery)Decorate(b Base){
	f.base = b
}

func (f Finery)Show(){
	if f.base!=nil{
		f.base.Show()
	}
}

type TShirts struct{
	Finery
}

func (t TShirts)Show(){
	fmt.Println("wear TShirts!!!")
	t.base.Show()
}

type BigTrouser struct{
	Finery
}

func (b BigTrouser)Show(){
	fmt.Println("wear big trouser!!!")
	b.base.Show()
}

func main(){
	xc := Person{name: "Cai"}

	fmt.Println("first collection: ")

	t := TShirts{Finery{}}
	big := BigTrouser{Finery{}}

	t.Decorate(xc)
	big.Decorate(t)

	big.Show()
}




