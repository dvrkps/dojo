package profile

import (
	"fmt"
	"testing"
)

func TestStrData(t *testing.T) {
	_ = StrData(10)
}

func BenchmarkStrData(b *testing.B) {
	var name string
	tests := []int{100, 1000, 10000, 100000}
	for _, tt := range tests {
		name = fmt.Sprintf("StrData(%d)", tt)

		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = StrData(tt)
			}
		})
	}
}
