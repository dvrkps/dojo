package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Date", "Tue, 07 Jun 2011 20:51:35 GMT")
	req.Header.Set("Authorization", `algorithm="hmac-sha256",headers="date",signature="UDysfR6MndUZReo07Y9r+vErn8vSxrnQ5ulit18iJ/Q=",apikey="_here_is_the_api_key_"`)
	w := httptest.NewRecorder()
	dummy := func(w http.ResponseWriter, r *http.Request) {}
	a := API{}
	handler := a.withAuth(dummy)
	handler(w, req)

	got := w.Result()

	const want = http.StatusOK
	if got.StatusCode != want {
		t.Errorf("withAuth status code = %v; want %v", got.StatusCode, want)
	}
}
