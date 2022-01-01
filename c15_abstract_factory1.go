package main

import "fmt"

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

func main() {
	u := User{
		id:   123,
		name: "byte",
	}

	su := SqlServerUser{}
	su.Insert(u)
	su.GetUser(u)
}
