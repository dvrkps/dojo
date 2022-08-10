package main

import (
	"fmt"
)

func main() {
	ur := &userRow{
		id:        1,
		firstName: "First",
		lastName:  "Last"}
	u, _ := do[User](ur)

	fmt.Printf("%#v\n\n", u)
}

type convertable[T any] interface {
	convert() T
}

func do[A any, B convertable[A]](r B) (A, error) {
	return r.convert(), nil
}
