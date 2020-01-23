package authorization

import (
	"net/http"
	"testing"
)

func TestValidate(t *testing.T) {
	for _, tt := range validateTests() {
		got := Validate(tt.header)
		fail := tt.fail
		t.Run(tt.name, func(t *testing.T) {
			if fail {
				if got == nil {
					t.Error("got nil; want error")
				}
				return
			}
			if got != nil {
				t.Errorf("got %v; want nil", got)
			}
		})
	}
}

type validateTest struct {
	name   string
	fail   bool
	header http.Header
}

func validateTests() []validateTest {
	tests := []validateTest{
		{
			name: "valid",
			fail: false,
			header: http.Header{
				"Authorization": []string{`algorithm="hmac-sha256",` +
					`headers="date",` +
					`signature="` +
					testSignatureValue +
					`",` +
					`apikey="_here_is_the_api_key_"`,
				},
				"Date": []string{"Tue, 07 Jun 2011 20:51:35 GMT"},
			},
		},
		{
			name: "no auth header",
			fail: true,
			header: http.Header{
				"Date": []string{"Tue, 07 Jun 2011 20:51:35 GMT"},
			},
		},
		{
			name:   "nil header",
			fail:   true,
			header: nil,
		},
	}

	return tests
}
