package main

import "fmt"

func main() {
	ur := &userRow{
		id:        1,
		firstName: "First",
		lastName:  "Last"}
	u, _ := do[User](ur)

	show("User", &u)

	pr := &productRow{
		id:    1,
		name:  "First",
		price: 1.23}
	p, _ := do[Product](pr)

	show("Product", &p)

}

type convertable[T any] interface {
	convert() T
}

func show[T fmt.Stringer](msg string, t T) {
	fmt.Printf("%s: %v\n", msg, t)
}

func do[A any, B convertable[A]](r B) (A, error) {
	return r.convert(), nil
}
