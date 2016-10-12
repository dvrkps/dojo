package profile

import (
	"fmt"
	"testing"
)

func TestData1(t *testing.T) {
	_ = Data1(10)
}

func TestStrData1(t *testing.T) {
	_ = StrData1(10)
}

func TestStrData2(t *testing.T) {
	_ = StrData2(10)
}

var benchs = []int{100} //, 1000, 10000}

func BenchmarkData(b *testing.B) {
	var name string
	for _, tt := range benchs {

		name = fmt.Sprintf("Data1(%d)", tt)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = Data1(tt)
			}
		})

		name = fmt.Sprintf("StrData1(%d)", tt)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = StrData1(tt)
			}
		})

		name = fmt.Sprintf("StrData2(%d)", tt)
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = StrData2(tt)
			}
		})

	}
}
