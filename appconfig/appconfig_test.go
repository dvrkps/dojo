package main

import (
	"log"
	"os"
	"testing"
)

var tests = []struct {
	cfg  *appConfig
	code int
}{
	{
		cfg: &appConfig{
			osargs: []string{"a", "b", "c"},
			stdout: nil,
			stderr: nil,
			logger: log.New(os.Stderr, "", 0),
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
			logger: log.New(os.Stderr, "", 0),
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
