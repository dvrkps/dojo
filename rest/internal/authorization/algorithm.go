package authorization

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

const algorithmHmacSha256 = "hmac-sha256"

type algorithm string

func (a *algorithm) check(message, signature string) error {
	if message == "" {
		return errors.New("empty message")
	}

	if signature == "" {
		return errors.New("empty signature")
	}

	switch *a {
	case algorithmHmacSha256:
		decoded, err := base64.StdEncoding.DecodeString(signature)
		if err != nil {
			return fmt.Errorf("%q: base64: %v", signature, err)
		}

		const sharedKey = "shared-key"

		mac := hmac.New(sha256.New, []byte(sharedKey))
		_, _ = mac.Write([]byte(message))
		expected := mac.Sum(nil)

		ok := hmac.Equal(decoded, expected)
		if !ok {
			return fmt.Errorf("%q: not equal", signature)
		}
	default:
		return fmt.Errorf("%q: invalid algorithm", *a)
	}

	return nil
}
