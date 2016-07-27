package main

import "testing"

func fakeAppConfig() *appConfig {
	return &appConfig{}
}

var tests = []struct {
	osargs []string
	code   int
}{
	{osargs: []string{"a", "b", "c"}, code: 0},
	{osargs: []string{}, code: 0},
}

func TestRunApp(t *testing.T) {
	for _, tt := range tests {
		ac := fakeAppConfig()
		ac.osargs = tt.osargs
		if got := runApp(ac); got != tt.code {
			t.Errorf("runApp(%v) = %d; want %d",
				tt.osargs, got, tt.code)
		}
	}
}
