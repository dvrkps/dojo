package authorization

import (
	"net/http"
	"testing"
)

func TestSignatureCheck(t *testing.T) {
	for _, tt := range signatureCheckTests() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.signature.check()
			if tt.fail {
				if err == nil {
					t.Error("fail error: got nil; want error")
				}
				return
			}
			if err != nil {
				t.Errorf("error: got %v; want nil", err)
			}
		})
	}
}

type signatureCheckTest struct {
	name      string
	fail      bool
	signature signature
}

const testSignatureValue = "druUFhOEGiAir3wspPxJF3mRU5CaXP0J7LIrN2tVDV8="

func signatureCheckTests() []signatureCheckTest {
	tests := []signatureCheckTest{
		{
			name: "valid",
			fail: false,
			signature: signature{
				value:     testSignatureValue,
				header:    http.Header{"Date": []string{"Tue, 07 Jun 2011 20:51:35 GMT"}},
				keys:      []string{"date"},
				algorithm: algorithm(algorithmHmacSha256),
			},
		},
		{
			name: "nil header",
			fail: true,
			signature: signature{
				value:     testSignatureValue,
				keys:      []string{"date"},
				algorithm: algorithm(algorithmHmacSha256),
			},
		},
		{
			name: "empty key fake",
			fail: true,
			signature: signature{
				value:     testSignatureValue,
				header:    http.Header{"Date": []string{"Tue, 07 Jun 2011 20:51:35 GMT"}},
				keys:      []string{"fake"},
				algorithm: algorithm(algorithmHmacSha256),
			},
		},
	}

	return tests
}
