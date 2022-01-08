package main

import "fmt"

var (
	hours        = []int{9, 10, 12, 13, 14, 17, 19, 22}
	Hour         = 0
	WorkFinished = false
)

func Work() {
	if Hour < 12 {
		fmt.Printf("Now: %d o'clock, moring, strong\n", Hour)
	} else if Hour < 13 {
		fmt.Printf("Now: %d o'clock, noon, hungry\n", Hour)
	} else if Hour < 17 {
		fmt.Printf("Now: %d o'clock, after noon, good\n", Hour)
	} else {
		if WorkFinished {
			fmt.Printf("Now: %d o'clock, go home\n", Hour)
		} else {
			if Hour < 21 {
				fmt.Printf("Now: %d o'clock, more work, tired\n", Hour)
			} else {
				fmt.Printf("Now: %d o'clock, sleeping\n", Hour)
			}
		}
	}
}

func main() {
	//fmt.Println("Please enter a number as Hour:")
	//fmt.Scan(&Hour)
	fmt.Println("Please enter yes/no as workfinished flag:")
	var flag string
	fmt.Scan(&flag)
	if flag == "yes" {
		WorkFinished = true
	}

	for _, h := range hours {
		Hour = h
		Work()
	}

}
