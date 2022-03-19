package main

import "fmt"

const exchangeRate = 7.53456

func main() {

	var tax itemFloat

	var it1, it2 itemFloat
	it1.SetEur(2.0, 20.00)
	it2.SetEur(4.0, 15.00)
	tax.SetHrk(2.0, 10.00)

	items := []itemFloat{it1, it2, tax}

	for _, i := range items {
		println(i.String())
	}

	allEur, allHrk := sumFloat(items)

	eur := calcEurFloat(allHrk)

	hrk := calcHrkFloat(allEur)

	if !isEqual(allEur, eur) {
		fmt.Printf("eur fail\n%f\n%f\n\n", allEur, eur)
	}

	if !isEqual(allHrk, hrk) {
		fmt.Printf("hrk fail\n%f\n%f\n\n", allHrk, hrk)
	}

}

func isEqual(a, b float64) bool {
	return a == b
}
