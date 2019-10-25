package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewCmdFlags(t *testing.T) {
	tests := []struct {
		fail bool
		args string
		want cmdFlags
	}{
		{
			fail: true,
			args: "cmd -m2 -ral",
		},
		{
			args: "cmd -m2 100 -ral 200 -chv 300",
			want: cmdFlags{
				m2:  100,
				ral: 200,
				chv: 300,
			},
		},
		{
			args: "cmd -version",
			want: cmdFlags{
				version: true,
			},
		},
	}

	for _, tt := range tests {
		args := strings.Fields(tt.args)
		var buf bytes.Buffer
		got, err := newCmdFlags(args, &buf)
		want := tt.want
		fail := tt.fail
		t.Run(tt.args, func(t *testing.T) {
			if fail {
				if err == nil {
					t.Errorf("err: got %v want <error>", err)
				}
				return
			}
			compareCmdFlags(t, got, want)
			if err != nil {
				t.Errorf("err: got %v want <nil>", err)
			}
		})
	}
}

func compareCmdFlags(t *testing.T, got, want cmdFlags) {
	if got.m2 != want.m2 {
		t.Errorf("m2: got %v; want %v", got.m2, want.m2)
	}
	if got.ral != want.ral {
		t.Errorf("ral: got %v; want %v", got.ral, want.ral)
	}
	if got.chv != want.chv {
		t.Errorf("chv: got %v; want %v", got.chv, want.chv)
	}
	if got.version != want.version {
		t.Errorf("version: got %v; want %v", got.version, want.version)
	}
}
