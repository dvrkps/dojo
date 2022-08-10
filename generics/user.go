package main

import "strconv"

type User struct {
	ID   string
	Name string
}

type userRow struct {
	firstName string
	lastName  string
	id        int
}

func (r *userRow) convert() User {
	return User{
		Name: r.firstName + " " + r.lastName,
		ID:   strconv.Itoa(r.id),
	}
}
