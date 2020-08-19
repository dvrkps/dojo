package main

import (
	"fmt"

	"github.com/dvrkps/dojo/monorepo/internal/website"
)

func main() {
	ws, err := website.All()
	if err != nil {
		fmt.Println(err)
	}

	for _, w := range ws {
		r, err := w.Parse(w)
		if err != nil {
			fmt.Printf("parse: %v\n", err)
			continue
		}
		fmt.Printf("%v: %v\n", w.Key, r)
	}
}
