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

type TestPaperA struct {
	test TestPaper
}

func (t TestPaperA) TestQuestion1() {
	t.test.TestQuestion1()
	fmt.Println("Answer: b")
}

func (t TestPaperA) TestQuestion2() {
	t.test.TestQuestion2()
	fmt.Println("Answer: b")
}

type TestPaperB struct {
	test TestPaper
}

func (t TestPaperB) TestQuestion1() {
	t.test.TestQuestion1()
	fmt.Println("Answer: b")
}

func (t TestPaperB) TestQuestion2() {
	t.test.TestQuestion2()
	fmt.Println("Answer: b")
}

func main() {
	fmt.Println("student A's test paper: ")
	studentA := TestPaperA{test: TestPaper{}}
	studentA.TestQuestion1()
	studentA.TestQuestion2()

	fmt.Println("student B's test paper: ")
	studentB := TestPaperB{TestPaper{}}
	studentB.TestQuestion1()
	studentB.TestQuestion2()
}
