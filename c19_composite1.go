package main

import "fmt"

type Company interface {
	Add(c Company)
	Remove(c Company)
	Display(depth int)
	LineOfDuty()
	SetName(s string)
	GetName() string
}

type CompanyStruct struct {
	name string
}

type ConcreteCompany struct {
	cs       CompanyStruct
	children []Company
}

func (c *ConcreteCompany) Add(company Company) {
	c.children = append(c.children, company)
}

func (c *ConcreteCompany) Remove(company Company) {
	index := 0
	for idx, value := range c.children {
		if value.GetName() == company.GetName() {
			index = idx
			break
		}
	}
	res := append(c.children[:index], c.children[index+1:]...)
	c.children = res
}

func (c *ConcreteCompany) Display(depth int) {
	pre := ""
	for i := 0; i < depth; i++ {
		pre += "-"
	}
	fmt.Printf("%s%s\n", pre, c.cs.name)
	for _, company := range c.children {
		company.Display(depth + 2)
	}
}

func (c *ConcreteCompany) SetName(s string) {
	c.cs.name = s
}

func (c *ConcreteCompany) GetName() string {
	return c.cs.name
}

func (c *ConcreteCompany) LineOfDuty() {
	for _, company := range c.children {
		company.LineOfDuty()
	}
}

type HRDepartment struct {
	cs CompanyStruct
}

func (h *HRDepartment) Add(c Company)    {}
func (h *HRDepartment) Remove(c Company) {}
func (h *HRDepartment) Display(depth int) {
	pre := ""
	for i := 0; i < depth; i++ {
		pre += "-"
	}
	fmt.Printf("%s%s\n", pre, h.cs.name)
}
func (h *HRDepartment) LineOfDuty()      { fmt.Println(h.GetName(), "Staff training") }
func (h *HRDepartment) SetName(s string) {}
func (h *HRDepartment) GetName() string  { return h.cs.name }

type FinanceDepartment struct {
	cs CompanyStruct
}

func (f *FinanceDepartment) Add(c Company)    {}
func (f *FinanceDepartment) Remove(c Company) {}
func (f *FinanceDepartment) Display(depth int) {
	pre := ""
	for i := 0; i < depth; i++ {
		pre += "-"
	}
	fmt.Printf("%s%s\n", pre, f.cs.name)
}
func (f *FinanceDepartment) LineOfDuty()      { fmt.Println(f.GetName(), "money, money, money!!!") }
func (f *FinanceDepartment) SetName(s string) {}
func (f *FinanceDepartment) GetName() string  { return f.cs.name }

func main() {
	root := &ConcreteCompany{cs: CompanyStruct{name: "???????????????"}}
	kk := &HRDepartment{cs: CompanyStruct{"????????????????????????"}}
	root.Add(kk)
	root.Add(&FinanceDepartment{cs: CompanyStruct{"??????????????????"}})

	comp := &ConcreteCompany{cs: CompanyStruct{name: "?????????????????????"}}
	comp.Add(&HRDepartment{cs: CompanyStruct{"??????????????????????????????"}})
	comp.Add(&FinanceDepartment{cs: CompanyStruct{"????????????????????????"}})
	root.Add(comp)

	comp1 := &ConcreteCompany{cs: CompanyStruct{name: "???????????????"}}
	comp1.Add(&HRDepartment{cs: CompanyStruct{"??????????????????????????????"}})
	comp1.Add(&FinanceDepartment{cs: CompanyStruct{"????????????????????????"}})
	root.Add(comp1)

	comp2 := &ConcreteCompany{cs: CompanyStruct{name: "???????????????"}}
	comp2.Add(&HRDepartment{cs: CompanyStruct{"??????????????????????????????"}})
	comp2.Add(&FinanceDepartment{cs: CompanyStruct{"????????????????????????"}})
	root.Add(comp2)

	fmt.Printf("\n????????????\n")
	root.Display(1)

	fmt.Printf("\n?????????")
	root.LineOfDuty()
}
