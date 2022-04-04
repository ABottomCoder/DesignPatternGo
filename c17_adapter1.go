package main

import "fmt"

type PlayerInterface interface {
	Attack()
	Defense()
}

type Player struct {
	name string
}

// forward
type Forwards struct {
	Player
}

func (f Forwards) Attack() {
	fmt.Printf("forward %s attack\n", f.name)
}

func (f Forwards) Defense() {
	fmt.Printf("forward %s defense\n", f.name)
}

// center
type Center struct {
	Player
}

func (c Center) Attack() {
	fmt.Printf("center %s attack\n", c.name)
}

func (c Center) Defense() {
	fmt.Printf("center %s defense\n", c.name)
}

// guards
type Guards struct {
	Player
}

func (g Guards) Attack() {
	fmt.Printf("guards %s attack\n", g.name)
}

func (g Guards) Defense() {
	fmt.Printf("guards %s defense\n", g.name)
}

func main() {
	b := Forwards{
		Player{
			name: "Battier",
		},
	}

	m := Guards{
		Player{
			name: "Middy",
		},
	}

	ym := Center{
		Player{
			name: "Yao",
		},
	}

	b.Attack()
	b.Defense()
	m.Attack()
	m.Defense()
	ym.Attack()
	ym.Defense()
}
