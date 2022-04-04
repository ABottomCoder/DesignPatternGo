package main

import "fmt"

type Observer interface {
	Update()
}

type ObserverStruct struct {
	name string
	sub  *Secretary
}

type Secretary struct {
	observers []Observer
	action    string
}

func (s *Secretary) Notify() {
	for _, observer := range s.observers {
		observer.Update()
	}
}

func (s *Secretary) SetAction(a string) {
	s.action = a
}

func (s *Secretary) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

type StockObserver struct {
	ObserverStruct
}

func (s *StockObserver) Update() {
	fmt.Printf("%s, %s close stock, work!!!\n", s.sub.action, s.name)
}

type NBAObserver struct {
	ObserverStruct
}

func (s *NBAObserver) Update() {
	fmt.Printf("%s, %s close NBA, work!!!\n", s.sub.action, s.name)
}

func main() {
	tongzizhe := &Secretary{}

	tongshi1 := &StockObserver{
		ObserverStruct{
			name: "Wei",
			sub:  tongzizhe,
		},
	}
	tongshi2 := &NBAObserver{
		ObserverStruct{
			name: "Yi",
			sub:  tongzizhe,
		},
	}

	tongzizhe.Attach(tongshi1)
	tongzizhe.Attach(tongshi2)

	tongzizhe.SetAction("boss is coming")
	tongzizhe.Notify()
}
