package profile

import (
	"fmt"
	"testing"
)

func TestStrData1(t *testing.T) {
	_ = strData1(10)
}

func BenchmarkStrData1(b *testing.B) {
	var name string
	tests := []int{100, 1000, 10000, 100000}
	for _, tt := range tests {
		name = fmt.Sprintf("strData1(%d)", tt)

		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = strData1(tt)
			}
		})
	}
}
