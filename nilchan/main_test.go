package main

import (
	"fmt"
	"reflect"
	"testing"
)

func wantValues(aName string, aMax int, bName string, bMax int) map[string]struct{} {
	all := make(map[string]struct{}, aMax+bMax)
	gen := func(all map[string]struct{}, name string, max int) {
		for i := 1; i <= max; i++ {
			k := fmt.Sprintf("%s%d", name, i)
			all[k] = struct{}{}
		}
	}
	gen(all, aName, aMax)
	gen(all, bName, bMax)
	return all
}

func Test(t *testing.T) {
	const (
		aName = "a"
		aMax  = 3
		bName = "b"
		bMax  = 5
	)
	want := wantValues(aName, aMax, bName, bMax)

	a := gen(aName, aMax)
	b := gen(bName, bMax)
	out := merge(a, b)
	got := make(map[string]struct{}, aMax+bMax)
	for o := range out {
		got[o] = struct{}{}
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("merge(%q:%d, %q:%d) = %v; want %v)",
			aName,
			aMax,
			bName,
			bMax,
			got,
			want,
		)
	}
}
