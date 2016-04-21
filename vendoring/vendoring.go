package main

import (
	"fmt"

	"github.com/dvrkps/dojo/vendoring/aa"
	"github.com/dvrkps/dojo/vendoring/bb"
)

func main() {

	fmt.Println("main")

	// ./aa will be ignored
	// ./vendor/aa is imported
	fmt.Println(aa.ID())

	// bb import ./bb/vendor/aa
	fmt.Println(bb.ID())

}
