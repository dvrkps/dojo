package iterbench

import "fmt"

func iterFor(max int) {
	var i int
	for {
		if ok := do(&i, &max); !ok {
			return
		}
	}
}

func do(i *int, lmt *int) bool {
	*i++
	fmt.Println(*i)
	return *i != *lmt
}
