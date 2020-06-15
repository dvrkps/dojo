package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	for _, tt := range testCases() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var (
				stdout bytes.Buffer
				stderr bytes.Buffer
			)
			args := strings.Fields(tt.args)
			gotExit := run(args, &stdout, &stderr)
			if gotExit != tt.exit {
				t.Errorf("exit: %v; want %v", gotExit, tt.exit)
			}
			if gotExit != exitOk {
				return
			}

			got := strings.TrimSpace(stdout.String())
			if got != tt.want {
				t.Errorf("result: %q; want %q", got, tt.want)
			}
		})
	}
}

type testCase struct {
	name string
	args string
	exit int
	want string
}

func testCases() []testCase {
	tests := []testCase{
		{
			name: "version",
			args: "cmd -version",
			exit: exitOk,
		},
		{
			name: "flag invalid",
			args: "cmd -invalid",
			exit: exitUser,
		},
		{
			name: "flag not a number",
			args: "cmd -m2 a",
			exit: exitUser,
		},
		{
			name: "flag < 0",
			args: "cmd -m2 -1",
			exit: exitUser,
		},
		{
			name: "empty flags",
			args: "cmd",
			exit: exitOk,
			want: "0 m2 = 0 ral, 0 chv",
		},
		{
			name: "1 m2",
			args: "cmd -m2 1",
			exit: exitOk,
			want: "1 m2 = 0 ral, 0 chv",
		},
		{
			name: "1 ral",
			args: "cmd -ral 1",
			exit: exitOk,
			want: "5754.542 m2 = 1 ral, 0 chv",
		},
		{
			name: "1 chv",
			args: "cmd -chv 1",
			exit: exitOk,
			want: "3.596652 m2 = 0 ral, 1 chv",
		},
		{
			name: "12121 m2",
			args: "cmd -m2 12121",
			exit: exitOk,
			want: "12121 m2 = 2 ral, 170 chv",
		},
	}

	return tests
}
