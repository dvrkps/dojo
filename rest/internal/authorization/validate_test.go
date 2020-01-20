package authorization

import (
	"net/http"
	"testing"
)

type validateTestType []struct {
	name   string
	fail   bool
	header http.Header
}

func validateTests() validateTestType {
	t := validateTestType{
		{
			name:   "nil header",
			fail:   true,
			header: nil,
		},
	}

	return t
}

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
