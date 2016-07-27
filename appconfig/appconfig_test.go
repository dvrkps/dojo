package main

import "testing"

var tests = []struct {
	osargs []string
	code   int
}{
	{osargs: []string{"a", "b", "c"}, code: 0},
	{osargs: []string{}, code: 0},
}

func fakeAppConfig() AppConfig {
	return AppConfig{}
}

func TestRealMain(t *testing.T) {
	ac := fakeAppConfig()
	for _, tt := range tests {
		ac.osargs = tt.osargs
		if got := realMain(ac); got != tt.code {
			t.Errorf("realMain(%v) = %d; want %d",
				tt.osargs, got, tt.code)
		}
	}
}
