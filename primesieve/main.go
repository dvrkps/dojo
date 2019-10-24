package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(run(50))
}

func run(max uint) string {
	var buf bytes.Buffer
	buf.WriteString(show("pure", pure(max)))
	buf.WriteString("\n")
	all := []kindPrimeser{
		newBase(max),
		newConcurrent(max),
	}

	for _, kp := range all {
		s := show(kp.kind(), kp.primes())
		buf.WriteString(s)
		buf.WriteString("\n")
	}
	return buf.String()
}

func show(kind string, primes []uint) string {
	const f = "%-11s: %d"
	return fmt.Sprintf(f, kind, primes)
}
