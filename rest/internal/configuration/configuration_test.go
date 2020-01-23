package configuration

import (
	"bytes"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	for _, tt := range newTests() {
		tt := tt
		t.Run(tt.args, func(t *testing.T) {
			args := strings.Fields(tt.args)

			var buf bytes.Buffer

			got, err := New(args, &buf)

			if tt.fail {
				if err == nil {
					t.Errorf("fail: got %v want <error>", err)
				}
				return
			}
			compareTestConfigurations(t, got, tt.want)
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

type newTest = struct {
	fail bool
	args string
	want *Configuration
}

func newTests() []newTest {
	tests := []newTest{
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

	return tests
}
