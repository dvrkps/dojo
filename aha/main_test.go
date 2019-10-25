package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "flag error",
			args: "cmd -m2 -ral",
			want: exitUser,
		},
		{
			name: "invalid m2",
			args: "cmd -m2 -1",
			want: exitErr,
		},
		{
			name: "empty flags",
			args: "cmd",
			want: exitOk,
		},
		{
			name: "valid",
			args: "cmd -m2 1",
			want: exitOk,
		},
		{
			name: "valid version",
			args: "cmd -m2 1 -version",
			want: exitOk,
		},
	}

	for _, tt := range tests {
		var buf bytes.Buffer
		args := strings.Fields(tt.args)
		got := run(args, &buf, &buf)
		want := tt.want
		t.Run(tt.name, func(t *testing.T) {
			if got != want {
				t.Errorf("exit = %v; want %v", got, want)

			}
		})
	}
}
