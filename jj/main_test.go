package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

type unOneTestCase struct {
	ok   bool
	in   []byte
	want interface{}
}

func unOneTestCases() map[string]unOneTestCase {
	c := map[string]unOneTestCase{
		"ok": {
			ok:   true,
			in:   []byte(`{"i":123, "f":123.456, "b":true,"s":"abc"}`),
			want: One{I: int64(123), F: float64(123.456), B: true, S: "abc"},
		},
	}
	return c
}

func TestUnmarshal(t *testing.T) {
	for name, tc := range unOneTestCases() {
		t.Run(name, func(t *testing.T) {
			testUnmarshal(t, tc)
		})
	}
}

func testUnmarshal(t *testing.T, tc unOneTestCase) {
	var got One
	err := json.Unmarshal(tc.in, &got)
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
