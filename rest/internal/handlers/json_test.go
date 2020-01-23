package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithJSON(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()

	dummy := func(w http.ResponseWriter, r *http.Request) {}

	a := API{}
	handler := a.withJSON(dummy)
	handler(w, r)

	res := w.Result()
	defer res.Body.Close()

	got := res.Header.Get(contentTypeHeader)

	if got != jsonContentType {
		t.Errorf("got %q; want %q", got, jsonContentType)
	}
}
