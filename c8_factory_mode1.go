package main

import "fmt"

type LeiFeng interface {
	Sweep()
	Wash()
	BuyRice()
}

type LearnLeiFeng struct {
}

func (l LearnLeiFeng) Sweep() {
	fmt.Println("Sweep!!!")
}

func (l LearnLeiFeng) Wash() {
	fmt.Println("Wash!!!")
}

func (l LearnLeiFeng) BuyRice() {
	fmt.Println("BuyRice!!!")
}

type UnderGraduate struct {
	LearnLeiFeng
}

type Volunteer struct {
	LearnLeiFeng
}

type IFactory interface {
	CreateLeiFeng()
}

type UndergraduateFactory struct {
}

func (u UndergraduateFactory) CreateLeiFeng() UnderGraduate {
	return UnderGraduate{}
}

type VolunteerFactory struct {
}

func (v VolunteerFactory) CreateLeiFeng() Volunteer {
	return Volunteer{}
}

func main() {
	factory := UndergraduateFactory{}
	undergraduate := factory.CreateLeiFeng()

	undergraduate.BuyRice()
	undergraduate.Sweep()
	undergraduate.Wash()

	return
}
