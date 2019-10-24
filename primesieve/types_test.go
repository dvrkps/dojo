package main

import "testing"

var testPrimeTypes = []struct {
	pt   primeType
	want string
}{
	{pt: 0, want: invalidTypeName},
	{pt: baseType, want: baseTypeName},
	{pt: concType, want: concTypeName},
}

func TestPrimeType_String(t *testing.T) {
	for _, tt := range testPrimeTypes {
		got := tt.pt.String()
		if got != tt.want {
			t.Errorf("PrimeType(%d) = %s; want %s", tt.pt, got, tt.want)
		}
	}

}

func testImplementsKindPrimeser(t *testing.T, i interface{}) {
	if _, ok := i.(kindPrimeser); !ok {
		t.Errorf("%T not implements kindPrimeser", i)
	}
}

func TestBaseGeneratorImplementsKindPrimeser(t *testing.T) {
	testImplementsKindPrimeser(t, &baseGenerator{})
}

func TestConcurrentGeneratorImplementsKindPrimeser(t *testing.T) {
	testImplementsKindPrimeser(t, &concurrentGenerator{})
}
