package main

import "fmt"

type User struct {
	Name string
	ID   int64
}

func (u *User) String() string {
	return fmt.Sprintf("%v: %v", u.ID, u.Name)
}

type userRow struct {
	firstName string
	lastName  string
	id        int64
}

func (r *userRow) convert() User {
	return User{
		Name: r.firstName + " " + r.lastName,
		ID:   r.id + 2000,
	}
}
