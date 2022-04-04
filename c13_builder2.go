package main

import "fmt"

type PersonDirector struct {
	builder PersonBuilder
}

func (p PersonDirector) Build() {
	p.builder.BuildHead()
	p.builder.BuildBody()
	p.builder.BuildLeftLeg()
	p.builder.BuildRightLeg()
	p.builder.BuildLeftArm()
	p.builder.BuildRightArm()
}

type PersonBuilder interface {
	BuildHead()
	BuildBody()
	BuildLeftLeg()
	BuildRightLeg()
	BuildLeftArm()
	BuildRightArm()
}

type PersonThinBuilder struct {
}

func (p PersonThinBuilder) BuildHead() {
	fmt.Println("draw a thin head")
}

func (p PersonThinBuilder) BuildBody() {
	fmt.Println("draw a thin Body")
}

func (p PersonThinBuilder) BuildLeftLeg() {
	fmt.Println("draw a thin left leg")
}

func (p PersonThinBuilder) BuildRightLeg() {
	fmt.Println("draw a thin right leg")
}

func (p PersonThinBuilder) BuildLeftArm() {
	fmt.Println("draw a thin left arm")
}

func (p PersonThinBuilder) BuildRightArm() {
	fmt.Println("draw a thin right arm")
}

func main() {
	director := PersonDirector{builder: PersonThinBuilder{}}
	director.Build()
}
