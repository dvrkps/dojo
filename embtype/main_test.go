package embtype

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNamedPoint(t *testing.T) {
	got := newNamedPoint("one", 1, 2)
	want := "one x:1 y:2"
	if got.String() != want {
		t.Fatalf("got %q; want %q", got, want)
	}

	gotJSON, err := json.Marshal(&got)
	if err != nil {
		t.Fatalf("marshall: %v", err)
	}

	wantJSON := []byte(`{"x":1,"y":2,"name":"one"}`)
	if !bytes.Equal(gotJSON, wantJSON) {
		t.Fatalf("json: got %s; want %s", gotJSON, wantJSON)
	}
}
