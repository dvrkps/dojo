package authorization

import (
	"testing"
)

func TestParseAuthField(t *testing.T) {
	tests := []struct {
		name  string
		fail  bool
		input string
		want  *authField
	}{
		{
			name: "valid",
			input: `"algorithm="hmac-sha256",` +
				`headers="date",` +
				`signature="UDysfR6MndUZReo07Y9r+vErn8vSxrnQ5ulit18iJ/Q=",` +
				`apikey="_here_is_the_api_key_"`,
			want: &authField{
				algorithm:  "hmac-sha256",
				headerKeys: []string{"date"},
				signature:  "UDysfR6MndUZReo07Y9r+vErn8vSxrnQ5ulit18iJ/Q=",
				apiKey:     "_here_is_the_api_key_",
			},
		},
		{
			name: "no algorithm",
			fail: true,
			input: `headers="date",` +
				`signature="UDysfR6MndUZReo07Y9r+vErn8vSxrnQ5ulit18iJ/Q=",` +
				`apikey="_here_is_the_api_key_"`,
		},
		{
			name: "no header keys",
			fail: true,
			input: `"algorithm="hmac-sha256",` +
				`signature="UDysfR6MndUZReo07Y9r+vErn8vSxrnQ5ulit18iJ/Q=",` +
				`apikey="_here_is_the_api_key_"`,
		},
		{
			name: "no signature",
			fail: true,
			input: `"algorithm="hmac-sha256",` +
				`headers="date",` +
				`apikey="_here_is_the_api_key_"`,
		},
		{
			name: "no api key",
			fail: true,
			input: `"algorithm="hmac-sha256",` +
				`headers="date",` +
				`signature="UDysfR6MndUZReo07Y9r+vErn8vSxrnQ5ulit18iJ/Q="`,
		},
		{
			name:  "empty input",
			fail:  true,
			input: "",
		},
	}
	for _, tt := range tests {
		got, err := parseAuthField(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			if tt.fail {
				if err == nil {
					t.Error("fail error: got nil; want error")
				}
				return
			}
			if err != nil {
				t.Errorf("error: got %v; want nil", err)
			}
			if got.testString() != tt.want.testString() {
				t.Errorf("got %v, want %v", got.testString(), tt.want.testString())
			}
		})
	}
}
