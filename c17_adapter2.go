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

// foreign center
type ForeignCenter struct {
	name string
}

func (f ForeignCenter) AttackChinese() {
	fmt.Printf("foreign center %s attack\n", f.name)
}

func (f ForeignCenter) DefenseChinese() {
	fmt.Printf("foreign center %s defense\n", f.name)
}

// translator
type Translator struct {
	Player
	ForeignCenter
}

func (t *Translator) SetForeignName() {
	t.ForeignCenter.name = t.Player.name
}

func (t Translator) Attack() {
	t.AttackChinese()
}

func (t Translator) Defense() {
	t.DefenseChinese()
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

	ym := Translator{
		Player{
			name: "Yao",
		},
		ForeignCenter{},
	}
	ym.SetForeignName()

	b.Attack()
	b.Defense()
	m.Attack()
	m.Defense()
	ym.Attack()
	ym.Defense()
}
