package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

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

func testUnmarshal(t *testing.T, tc testUnmarshalCase) {
	var got T
	err := json.Unmarshal(tc.in, got)
	if !tc.ok {
		if !reflect.DeepEqual(got, tc.want) || err == nil {
			t.Errorf("got %+v, %v; want %+v, <error>", got, err, tc.want)
		}
		return
	}
	if !reflect.DeepEqual(got, tc.want) || err != nil {
		t.Errorf("got %+v, %+v; want %+v, <nil>", got, err, tc.want)
	}
}
