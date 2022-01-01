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
	fmt.Printf("search id %d int Access table\n", u.id)
}

type IFactory interface {
	CreateUser() IUser
}

type SqlServerFactory struct {
}

func (s SqlServerFactory) CreateUser() IUser {
	return SqlServerUser{}
}

type AccessServerFactory struct {
}

func (a AccessServerFactory) CreateUser() IUser {
	return AccessUser{}
}

func main() {
	u := User{}

	factory := SqlServerFactory{}

	iu := factory.CreateUser()

	iu.Insert(u)
	iu.GetUser(u)
}
