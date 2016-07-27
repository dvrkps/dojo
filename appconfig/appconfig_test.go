package main

import (
	"io"
	"testing"
)

func fakeAppConfig(osargs []string, stdout, stderr io.Writer) *appConfig {
	if len(osargs) == 0 && stdout == nil && stderr == nil {
		return nil
	}
	return &appConfig{
		osargs: osargs,
		stdout: stdout,
		stderr: stderr}
}

var tests = []struct {
	cfg  *appConfig
	code int
}{
	{
		cfg: &appConfig{
			osargs: []string{"a", "b", "c"},
			stdout: nil,
			stderr: nil,
		},
		code: 0,
	},
	{
		cfg:  nil,
		code: 1,
	},
	{
		cfg: &appConfig{
			osargs: []string{},
			stdout: nil,
			stderr: nil,
		},
		code: 1,
	},
}

func TestRunApp(t *testing.T) {
	for _, tt := range tests {
		if got := runApp(tt.cfg); got != tt.code {
			t.Errorf("runApp(%v) = %d; want %d",
				tt.cfg, got, tt.code)
		}
	}
}
