package main

import "fmt"

type ICloneable interface {
	Clone() ICloneable
}

type WorkExperience struct {
	timeArea string
	company  string
}

func (w WorkExperience) Clone() ICloneable {
	return w
}

type Resume struct {
	name           string
	sex            string
	age            string
	workExperience *WorkExperience
}

func NewResume(name string) Resume {
	return Resume{
		name:           name,
		workExperience: &WorkExperience{},
	}
}

func (r *Resume) SetPersonalInfo(sex string, age string) {
	r.sex = sex
	r.age = age
}

func (r *Resume) SetWorkExperience(timeArea string, company string) {
	r.workExperience.timeArea = timeArea
	r.workExperience.company = company
}

func (r *Resume) Display() {
	fmt.Printf("%s %s %s\n", r.name, r.sex, r.age)
	fmt.Printf("work experience: %s %s\n", r.workExperience.timeArea, r.workExperience.company)
}

func (r Resume) Clone() ICloneable {
	copyWorkExperience := r.workExperience.Clone().(WorkExperience)
	r.workExperience = &copyWorkExperience
	return r
}

func main() {
	resumeA := NewResume("BigBird")
	resumeA.SetPersonalInfo("male", "29")
	resumeA.SetWorkExperience("1998-2000", "bytedance")

	resumeB := resumeA.Clone().(Resume)
	resumeB.SetPersonalInfo("female", "33")
	resumeB.SetWorkExperience("2000-2010", "CQ")

	resumeA.Display()
	resumeB.Display()
}
