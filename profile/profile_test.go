package profile

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	_ = Data(10)
}

func TestStrData(t *testing.T) {
	_ = StrData(10)
}

var benchs = []int{100} //, 1000, 10000}

func BenchmarkData(b *testing.B) {
	var name string
	for _, tt := range benchs {

		name = fmt.Sprintf("Data(%d)", tt)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = Data(tt)
			}
		})

		name = fmt.Sprintf("StrData(%d)", tt)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = StrData(tt)
			}
		})

	}
}
