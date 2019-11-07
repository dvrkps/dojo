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
		want Configuration
	}{
		{
			args: "cmd",
			want: Configuration{
				Debug: defaultFlagDebug,
				Port:  defaultFlagPort,
			},
		},
		{
			args: "cmd -debug -port 6666",
			want: Configuration{
				Debug: true,
				Port:  6666,
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

func compareTestConfigurations(t *testing.T, got, want Configuration) {
	if got.Debug != want.Debug {
		t.Errorf("debug: got %v; want %v", got.Debug, want.Debug)
	}

	if got.Port != want.Port {
		t.Errorf("port: got %v; want %v", got.Port, want.Port)
	}
}
