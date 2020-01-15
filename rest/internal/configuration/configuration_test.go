package configuration

import (
	"bytes"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		fail bool
		args string
		want *Configuration
	}{
		{
			args: "cmd",
			want: &Configuration{
				Addr:    defaultFlagAddr,
				Verbose: defaultFlagVerbose,
			},
		},
		{
			args: "cmd -addr localhost:8080 -v",
			want: &Configuration{
				Addr:    defaultFlagAddr,
				Verbose: true,
			},
		},
	}

	for _, tt := range tests {
		args := strings.Fields(tt.args)

		var buf bytes.Buffer

		got, err := New(args, &buf)
		want := tt.want
		fail := tt.fail
		t.Run(tt.args, func(t *testing.T) {
			if fail {
				if err == nil {
					t.Errorf("fail: got %v want <error>", err)
				}
				return
			}
			compareTestConfigurations(t, got, want)
			if err != nil {
				t.Errorf("compare: got %v want <nil>", err)
			}
		})
	}
}

func compareTestConfigurations(t *testing.T, got, want *Configuration) {
	if got.Addr != want.Addr {
		t.Errorf("addr: got %v; want %v", got.Addr, want.Addr)
	}

	if got.Verbose != want.Verbose {
		t.Errorf("verbose: got %v; want %v", got.Verbose, want.Verbose)
	}
}
