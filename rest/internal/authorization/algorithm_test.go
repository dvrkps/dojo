package authorization

import "testing"

func TestAlgorithmCheck(t *testing.T) {
	for _, tt := range algorithmCheckTests() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := tt.algorithm.check(tt.message, tt.signature)
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

type algorithmCheckTest struct {
	name      string
	fail      bool
	algorithm algorithm
	signature string
	message   string
}

const testMessage = "title\ndate: Tue, 07 Jun 2011 20:51:35 GMT"

func algorithmCheckTests() []algorithmCheckTest {
	tests := []algorithmCheckTest{
		{
			name:      "valid hmac-sha256",
			fail:      false,
			algorithm: algorithm(algorithmHmacSha256),
			signature: testSignatureValue,
			message:   testMessage,
		},
		{
			name:      "invalid algorithm",
			fail:      true,
			algorithm: algorithm("invalid"),
			signature: testSignatureValue,
			message:   testMessage,
		},
		{
			name:      "base64 decode hmac-sha256",
			fail:      true,
			algorithm: algorithm(algorithmHmacSha256),
			signature: "a bc",
			message:   testMessage,
		},
		{
			name:      "not equal hmac-sha256",
			fail:      true,
			algorithm: algorithm("hmac-sha256"),
			signature: "Uabc",
			message:   testMessage,
		},
		{
			name: "empty message",
			fail: true,
		},
		{
			name:      "empty signature",
			fail:      true,
			signature: "",
			message:   testMessage,
		},
	}

	return tests
}
