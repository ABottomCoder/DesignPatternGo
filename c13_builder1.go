package main

import "fmt"

type PersonThinBuilder struct {
}

func (p PersonThinBuilder) DrawHead() {
	fmt.Println("draw a head")
}

func (p PersonThinBuilder) DrawBody() {
	fmt.Println("draw a Body")
}

func (p PersonThinBuilder) DrawLeftLeg() {
	fmt.Println("draw a left leg")
}

func (p PersonThinBuilder) DrawRightLeg() {
	fmt.Println("draw a right leg")
}

func (p PersonThinBuilder) DrawLeftArm() {
	fmt.Println("draw a left arm")
}

func (p PersonThinBuilder) DrawRightArm() {
	fmt.Println("draw a right arm")
}

func main() {
	p := PersonThinBuilder{}
	p.DrawHead()
	p.DrawBody()
	p.DrawLeftLeg()
	p.DrawRightLeg()
	p.DrawLeftArm()
	p.DrawRightArm()
}
