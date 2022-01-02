package json

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	in := struct {
		ID int64 `json:"id,string"`
	}{
		ID: 42,
	}

	const want = `{"id":"42"}`
	testMarshal(t, in, want)
}

func testMarshal(t *testing.T, in interface{}, want string) {
	t.Helper()

	got, err := json.Marshal(&in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	if string(got) != want {
		t.Errorf("got %s; want %s", got, want)
	}
}
