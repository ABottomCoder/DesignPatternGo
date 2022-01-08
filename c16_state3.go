package main

import "fmt"

type Work struct {
	current StateInterface
	hour    int
	finish  bool
}

func (w *Work) WriteProgram() {
	w.current.WriteProgram(w)
}

type StateInterface interface {
	WriteProgram(w *Work)
}

// forenoon state
type ForeNoonState struct {
}

func (f ForeNoonState) WriteProgram(w *Work) {
	if w.hour < 12 {
		fmt.Printf("Now: %d o'clock, moring, strong\n", w.hour)
	} else {
		w.current = NoonState{}
		w.WriteProgram()
	}
}

// noon state
type NoonState struct {
}

func (n NoonState) WriteProgram(w *Work) {
	if w.hour < 13 {
		fmt.Printf("Now: %d o'clock, noon, hungry\n", w.hour)
	} else {
		w.current = AfterNoonState{}
		w.WriteProgram()
	}
}

// afternoon state
type AfterNoonState struct {
}

func (a AfterNoonState) WriteProgram(w *Work) {
	if w.hour < 17 {
		fmt.Printf("Now: %d o'clock, after noon, good\n", w.hour)
	} else {
		w.current = EveningState{}
		w.WriteProgram()
	}
}

// evening state
type EveningState struct {
}

func (a EveningState) WriteProgram(w *Work) {
	if w.finish {
		w.current = RestState{}
		w.WriteProgram()
	} else {
		if w.hour < 21 {
			fmt.Printf("Now: %d o'clock, more work, tired\n", w.hour)
		} else {
			w.current = RestState{}
			w.WriteProgram()
		}
	}
}

// sleeping state
type SleepingState struct {
}

func (s SleepingState) WriteProgram(w *Work) {
	fmt.Printf("Now: %d o'clock, sleeping\n", w.hour)
}

// rest state
type RestState struct {
}

func (r RestState) WriteProgram(w *Work) {
	fmt.Printf("Now: %d o'clock, go home\n", w.hour)
}

var hours = []int{9, 10, 12, 13, 14, 17, 19, 22}

func main() {
	w := Work{
		current: ForeNoonState{},
	}

	fmt.Println("Please enter yes/no as workfinished flag:")
	var flag string
	fmt.Scan(&flag)
	if flag == "yes" {
		w.finish = true
	}

	for _, h := range hours {
		w.hour = h
		w.WriteProgram()
	}
}
