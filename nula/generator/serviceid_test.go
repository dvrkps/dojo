package main

import (
	"os"
	"strconv"
	"testing"
)

func TestEnvServiceID(t *testing.T) {
	want := os.Getenv(keyServiceID)
	if got := EnvServiceID(); got != want {
		t.Errorf("EnvServiceID() = %q; want %q",
			got, want)
	}
}

func TestServiceID(t *testing.T) {

	min := strconv.Itoa(minServiceID)
	max := strconv.Itoa(maxServiceID)

	tests := []struct {
		sid  string
		want int
	}{
		{sid: "42", want: 42},
		{sid: min, want: minServiceID},
		{sid: max, want: maxServiceID},
	}

	for _, tt := range tests {
		got, err := ServiceID(tt.sid)
		if got != tt.want || err != nil {
			t.Errorf("ServiceID(%q) = %d, %v; want %d, <nil>",
				tt.sid, got, err, tt.want)
		}
	}
}

func TestServiceID_errors(t *testing.T) {

	min := strconv.Itoa(minServiceID - 1)
	max := strconv.Itoa(maxServiceID + 1)

	tests := []struct {
		sid string
	}{
		{sid: ""},
		{sid: "NaN"},
		{sid: min},
		{sid: max},
	}

	const want = 0
	for _, tt := range tests {
		got, err := ServiceID(tt.sid)
		if got != want || err == nil {
			t.Errorf("ServiceID(%q) = %d, %v; want %d, <error>",
				tt.sid, got, err, want)
		}
	}
}
