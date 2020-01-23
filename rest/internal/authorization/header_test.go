package authorization

import (
	"testing"
)

func TestNewHeader(t *testing.T) {
	for _, tt := range newHeaderTests() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := newHeader(tt.input)
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

type newHeaderTest struct {
	name  string
	fail  bool
	input string
	want  *header
}

func newHeaderTests() []newHeaderTest {
	tests := []newHeaderTest{
		{
			name: "valid",
			input: `"algorithm="hmac-sha256",` +
				`headers="date",` +
				`signature="` +
				testSignatureValue +
				`",` +
				`apikey="_here_is_the_api_key_"`,
			want: &header{
				algorithm: algorithm(algorithmHmacSha256),
				keys:      []string{"date"},
				signature: testSignatureValue,
				apiKey:    "_here_is_the_api_key_",
			},
		},
		{
			name:  "empty input",
			fail:  true,
			input: "",
		},
	}

	return tests
}
