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
		r, err := w.Index(w)
		if err != nil {
			fmt.Printf("index: %v\n", err)
			continue
		}
		fmt.Printf("%v: %v\n", w.Key, r)
	}
}
