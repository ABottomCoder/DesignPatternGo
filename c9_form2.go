package main

import "fmt"

type ICloneable interface {
	Clone() ICloneable
}

type Resume struct {
	name     string
	sex      string
	age      string
	timeArea string
	company  string
}

func NewResume(name string) *Resume {
	return &Resume{
		name: name,
	}
}

func (r *Resume) SetPersonalInfo(sex string, age string) {
	r.sex = sex
	r.age = age
}

func (r *Resume) SetWorkExperience(timeArea string, company string) {
	r.timeArea = timeArea
	r.company = company
}

func (r *Resume) Display() {
	fmt.Printf("%s %s %s\n", r.name, r.sex, r.age)
	fmt.Printf("work experience: %s %s\n", r.timeArea, r.company)
}

func (r Resume) Clone() ICloneable {
	return r
}

func main() {
	resumeA := NewResume("BigBird")
	resumeA.SetPersonalInfo("male", "29")
	resumeA.SetWorkExperience("1998-2000", "bytedance")
	resumeA.Display()

	resumeB := resumeA.Clone().(Resume)
	resumeB.SetPersonalInfo("female", "33")
	resumeB.Display()
}
