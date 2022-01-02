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

	got, err := json.Marshal(&in)
	if err != nil {
		t.Error(err)
	}

	const want = `{"id":"42"}`
	if string(got) != want {
		t.Errorf("got %s; want %s", got, want)
	}
}
