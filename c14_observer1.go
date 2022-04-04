package main

import "fmt"

type Secretary struct {
	observers []*StockObserver
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

func (s *Secretary) Attach(observer *StockObserver) {
	s.observers = append(s.observers, observer)
}

type StockObserver struct {
	name string
	sub  *Secretary
}

func (s *StockObserver) Update() {
	fmt.Printf("%s, %s close stock, work!!!\n", s.sub.action, s.name)
}

func NewStockObserver(name string, sub *Secretary) *StockObserver {
	return &StockObserver{
		name: name,
		sub:  sub,
	}
}

func main() {
	tongzizhe := &Secretary{}

	tongshi1 := NewStockObserver("Wei", tongzizhe)
	tongshi2 := NewStockObserver("Yi", tongzizhe)

	tongzizhe.Attach(tongshi1)
	tongzizhe.Attach(tongshi2)

	tongzizhe.SetAction("boss is coming")
	tongzizhe.Notify()
}
