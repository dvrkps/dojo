package reader

import "testing"

func TestRead(t *testing.T) {
	content := "abcdefghijk"

	err := read(content)
	if err != nil {
		t.Log(err)
	}
}
