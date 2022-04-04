package main

import "fmt"

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

func main() {
	// before battle
	lixiaoyao := GameRole{}
	lixiaoyao.GetInitState()
	fmt.Println("before battle")
	lixiaoyao.StateDisplay()

	// save status
	backup := GameRole{}
	backup.atk = lixiaoyao.atk
	backup.def = lixiaoyao.def
	backup.vit = lixiaoyao.vit

	// battle
	lixiaoyao.Fight()
	fmt.Println("after battle")
	lixiaoyao.StateDisplay()

	// backup
	lixiaoyao.atk = backup.atk
	lixiaoyao.def = backup.def
	lixiaoyao.vit = backup.vit
	fmt.Println("after backup")
	lixiaoyao.StateDisplay()
}
