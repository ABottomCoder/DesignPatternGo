package main

import "fmt"

var hours = []int{9, 10, 12, 13, 14, 17, 19, 22}

type Work struct {
	hour   int
	finish bool
}

func (w Work) WriteProgram() {
	if w.hour < 12 {
		fmt.Printf("Now: %d o'clock, moring, strong\n", w.hour)
	} else if w.hour < 13 {
		fmt.Printf("Now: %d o'clock, noon, hungry\n", w.hour)
	} else if w.hour < 17 {
		fmt.Printf("Now: %d o'clock, after noon, good\n", w.hour)
	} else {
		if w.finish {
			fmt.Printf("Now: %d o'clock, go home\n", w.hour)
		} else {
			if w.hour < 21 {
				fmt.Printf("Now: %d o'clock, more work, tired\n", w.hour)
			} else {
				fmt.Printf("Now: %d o'clock, sleeping\n", w.hour)
			}
		}
	}
}

func main() {
	w := Work{}

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
