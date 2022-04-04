package main

import "fmt"

type TestPaper struct {
}

func (t TestPaper) TestQuestion1() {
	fmt.Println("This is question 1")
}

func (t TestPaper) TestQuestion2() {
	fmt.Println("This is question 2")
}

func (t TestPaper) Answer1() {
}

func (t TestPaper) Answer2() {
}

type TestPaperA struct {
	test TestPaper
}

func (t TestPaperA) Answer1() {
	t.test.TestQuestion1()
	fmt.Println("Answer: b")
}

func (t TestPaperA) Answer2() {
	t.test.TestQuestion2()
	fmt.Println("Answer: b")
}

type TestPaperB struct {
	test TestPaper
}

func (t TestPaperB) Answer1() {
	t.test.TestQuestion1()
	fmt.Println("Answer: b")
}

func (t TestPaperB) Answer2() {
	t.test.TestQuestion2()
	fmt.Println("Answer: b")
}

func main() {
	fmt.Println("student A's test paper: ")
	studentA := TestPaperA{test: TestPaper{}}
	studentA.Answer1()
	studentA.Answer2()

	fmt.Println("student B's test paper: ")
	studentB := TestPaperB{TestPaper{}}
	studentB.Answer1()
	studentB.Answer2()
}
