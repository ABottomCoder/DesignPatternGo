package main

import "fmt"

// box of role state
type RoleStateMemento struct {
	vit int
	atk int
	def int
}

func (r *RoleStateMemento) Init(vit, atk, def int) {
	r.vit = vit
	r.atk = atk
	r.def = def
}

type RoleStateCaretaker struct {
	memento RoleStateMemento
}

func (r *RoleStateCaretaker) Get() RoleStateMemento {
	return r.memento
}
func (r *RoleStateCaretaker) Set(m RoleStateMemento) {
	r.memento = m
}

type GameRole struct {
	vit int
	atk int
	def int
}

func (g *GameRole) StateDisplay() {
	fmt.Println("current role state:")
	fmt.Printf("hp: %d\n", g.vit)
	fmt.Printf("atk: %d\n", g.atk)
	fmt.Printf("def: %d\n", g.def)
}

func (g *GameRole) GetInitState() {
	g.vit = 100
	g.atk = 100
	g.def = 100
}

func (g *GameRole) Fight() {
	g.vit = 0
	g.atk = 0
	g.def = 0
}

func (g *GameRole) SaveState() RoleStateMemento {
	return RoleStateMemento{
		vit: g.vit,
		atk: g.atk,
		def: g.def,
	}
}

func (g *GameRole) RecoveryState(memento RoleStateMemento) {
	g.vit = memento.vit
	g.def = memento.def
	g.atk = memento.atk
}

func main() {
	// before battle
	fmt.Println("before battle...")
	lixiaoyao := GameRole{}
	lixiaoyao.GetInitState()
	lixiaoyao.StateDisplay()
	fmt.Println("")

	// save state
	fmt.Println("save state...")
	stateAdmin := RoleStateCaretaker{}
	stateAdmin.Set(lixiaoyao.SaveState())
	fmt.Println("")

	// battle
	fmt.Println("battle...")
	lixiaoyao.Fight()
	lixiaoyao.StateDisplay()
	fmt.Println("")

	// recover
	fmt.Println("recover...")
	lixiaoyao.RecoveryState(stateAdmin.memento)
	lixiaoyao.StateDisplay()

}
