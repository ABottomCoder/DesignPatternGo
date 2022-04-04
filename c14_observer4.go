package main

import "fmt"

type ObserverStruct struct {
	name string
	sub  Subject
}

type Subject interface {
	Notify()
	GetAction() string
}

type SubjectStruct struct {
	action string
	fs     []func()
}

func (s *SubjectStruct) Notify() {
	for _, f := range s.fs {
		f()
	}
}

func (s *SubjectStruct) GetAction() string {
	return s.action
}

type Boss struct {
	SubjectStruct
}

type Secretary struct {
	SubjectStruct
}

type StockObserver struct {
	ObserverStruct
}

func (s *StockObserver) CloseStockMarket() {
	fmt.Printf("%s, %s close stock, work!!!\n", s.sub.GetAction(), s.name)
}

type NBAObserver struct {
	ObserverStruct
}

func (s *NBAObserver) CloseNBADirectSeeding() {
	fmt.Printf("%s, %s close NBA, work!!!\n", s.sub.GetAction(), s.name)
}

func main() {
	boss := &Boss{
		SubjectStruct{
			action: "I'm boss",
		},
	}

	tongshi1 := &StockObserver{
		ObserverStruct{
			name: "Wei",
			sub:  boss,
		},
	}
	tongshi2 := &NBAObserver{
		ObserverStruct{
			name: "Yi",
			sub:  boss,
		},
	}

	boss.fs = append(boss.fs, tongshi1.CloseStockMarket)
	boss.fs = append(boss.fs, tongshi2.CloseNBADirectSeeding)
	boss.Notify()
}
