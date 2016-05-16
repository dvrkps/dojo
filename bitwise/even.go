package main

import (
	"fmt"
	"os"
)

func even() {

	vals := [...]struct {
		n    int
		want bool
	}{
		{n: 2, want: true},
		{n: 5, want: false},
	}

	for _, v := range vals {
		bit := isEvenBitwise(v.n)
		rem := isEvenRemainder(v.n)
		if bit != rem {
			fmt.Println("bitwise != remainder")
			os.Exit(1)
		}
		if got, want := bit, v.want; got != want {
			fmt.Printf("isEvenBitwise(%d) = %v; want %v\n",
				v.n, got, want)
		}
		if got, want := rem, v.want; got != want {
			fmt.Printf("isEvenRemainder(%d) = %v; want %v\n",
				v.n, got, want)
		}
	}

}

func isEvenBitwise(i int) bool {
	v := (i & 1) == 0
	return v
}

func isEvenRemainder(i int) bool {
	v := (i % 2) == 0
	return v
}
