package main

import "fmt"

type IUser interface {
	Insert(u User)
	GetUser(u User)
}

type User struct {
	id   int
	name string
}

type SqlServerUser struct {
}

func (s SqlServerUser) Insert(u User) {
	fmt.Printf("insert user %s to SQL table\n", u.name)
}

func (s SqlServerUser) GetUser(u User) {
	fmt.Printf("search id %d int SQL table\n", u.id)
}

type AccessUser struct {
}

func (a AccessUser) Insert(u User) {
	fmt.Printf("insert user %s to Access table\n", u.name)
}

func (a AccessUser) GetUser(u User) {
	fmt.Printf("search user %d int Access table\n", u.id)
}

type IFactory interface {
	CreateUser() IUser
	CreateDepartment() IDepartment
}

type SqlServerFactory struct {
}

func (s SqlServerFactory) CreateUser() IUser {
	return SqlServerUser{}
}

func (s SqlServerFactory) CreateDepartment() IDepartment {
	return SqlServerDepartment{}
}

type AccessServerFactory struct {
}

func (a AccessServerFactory) CreateUser() IUser {
	return AccessUser{}
}

func (a AccessServerFactory) CreateDepartment() IDepartment {
	return AccessServerDepartment{}
}

type Department struct {
	id         int
	department string
}

type IDepartment interface {
	Insert(d Department)
	GetDepartment(d Department)
}

type SqlServerDepartment struct {
}

func (s SqlServerDepartment) Insert(d Department) {
	fmt.Printf("insert department %s to SQL table\n", d.department)
}

func (s SqlServerDepartment) GetDepartment(d Department) {
	fmt.Printf("search department %d int SQL table\n", d.id)
}

type AccessServerDepartment struct {
}

func (s AccessServerDepartment) Insert(d Department) {
	fmt.Printf("insert department %s to Access table\n", d.department)
}

func (s AccessServerDepartment) GetDepartment(d Department) {
	fmt.Printf("search department %d int Access table\n", d.id)
}

func main() {
	u := User{}
	d := Department{}

	factory := AccessServerFactory{}

	iu := factory.CreateUser()

	iu.Insert(u)
	iu.GetUser(u)

	id := factory.CreateDepartment()
	id.Insert(d)
	id.GetDepartment(d)
}
