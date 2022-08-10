package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	ID    int64
}

func (p *Product) String() string {
	return fmt.Sprintf("%v: %v %v", p.ID, p.Name, p.Price)
}

type productRow struct {
	name  string
	price float64
	id    int64
}

func (r *productRow) convert() Product {
	return Product{
		Name:  "product " + r.name,
		ID:    r.id + 1000,
		Price: r.price + 100,
	}
}
