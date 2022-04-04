package main

import "fmt"

type Observer interface {
	Update()
}

type ObserverStruct struct {
	name string
	sub  Subject
}

type Subject interface {
	Attach(ob Observer)
	Detach(ob Observer)
	Notify()
	GetAction() string
}

type SubjectStruct struct {
	observers []Observer
	action    string
}

func (s *SubjectStruct) Attach(ob Observer) {
	s.observers = append(s.observers, ob)
}

func (s *SubjectStruct) Detach(ob Observer) {
	var index = -1
	for idx, observer := range s.observers {
		if observer == ob {
			index = idx
			break
		}
	}
	if index == -1 {
		return
	} else if index == 0 {
		s.observers = s.observers[1:]
	} else if index == len(s.observers)-1 {
		s.observers = s.observers[:index-1]
	} else {
		s.observers = append(s.observers[:index-1], s.observers[index:]...)
	}
}

func (s *SubjectStruct) Notify() {
	for _, observer := range s.observers {
		observer.Update()
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

func (s *StockObserver) Update() {
	fmt.Printf("%s, %s close stock, work!!!\n", s.sub.GetAction(), s.name)
}

type NBAObserver struct {
	ObserverStruct
}

func (s *NBAObserver) Update() {
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

	boss.Attach(tongshi1)
	boss.Attach(tongshi2)
	boss.Notify()

	boss.Detach(tongshi1)
	boss.Notify()
}
