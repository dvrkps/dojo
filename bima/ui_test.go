package main

import (
	"strings"
	"testing"
)

func TestPrintRow(t *testing.T) {
	test := func(h, bmi float64, want string) {
		trim := strings.TrimSpace
		if got := printRow(h, bmi); got != want {
			t.Errorf("printRow(%v,%v) = %v; want %v",
				h, bmi, trim(got), trim(want))
		}
	}
	test(1.87, 20, "20.0 (69.94 kg)\n")
	test(1.87, 22, "22.0 (76.93 kg)\n")
	test(1.87, 0, " 0.0 (0.00 kg)\n")
}
