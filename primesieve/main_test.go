package main

import (
	"reflect"
	"testing"
)

func testPrimes(max uint) []uint {
	exists := map[uint]struct{}{}
	primes := []uint{}
	for i := uint(2); i < max; i++ {
		if _, ok := exists[i]; ok {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < max; k += i {
			exists[k] = struct{}{}
		}
	}
	return primes
}

const testMax = 5

var primesTestCases = []struct {
	kp kindPrimeser
}{
	{kp: newBase(testMax)},
	{kp: newConcurrent(testMax)},
}

func TestPrimes(t *testing.T) {
	want := testPrimes(testMax)
	for _, tt := range primesTestCases {
		got := tt.kp.primes()
		name := tt.kp.kind()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s(%v) = %v; want %v", name, testMax, got, want)
		}
	}
}

func TestRun(t *testing.T) {
	got := run(5)
	want := "pure       : [2 3]\n" +
		"base       : [2 3]\n" +
		"concurrent : [2 3]\n"
	if got != want {
		t.Errorf("run(5)\ngot:\n%s\nwant:\n%s", got, want)
	}
}
