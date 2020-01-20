package authorization

import (
	"net/http"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name   string
		fail   bool
		header http.Header
	}{
		{
			name:   "nil header",
			fail:   true,
			header: nil,
		},
	}
	for _, tt := range tests {
		got := Validate(tt.header)
		t.Run(tt.name, func(t *testing.T) {
			if tt.fail {
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
