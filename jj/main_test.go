package main

import "testing"

type testUnmarshalCase struct {
	ok   bool
	in   []byte
	want interface{}
}

func testUnmarshalCases() map[string]testUnmarshalCase {
	c := map[string]testUnmarshalCase{
		"ok": {
			ok:   true,
			in:   []byte(`{i:123, f: 123.456, b:true,s:"abc"}`),
			want: T{I: int64(123), F: float64(123.456), B: true, S: []byte{'a', 'b', 'c'}},
		},
	}
	return c
}

func TestUnmarshal(t *testing.T) {}
