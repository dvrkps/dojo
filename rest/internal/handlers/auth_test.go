package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	dummy := func(w http.ResponseWriter, r *http.Request) {}
	a := API{}
	handler := jsonContentType(a.withAuth(dummy))
	handler(w, req)

	got := w.Result()

	fmt.Println(got.StatusCode)
	fmt.Println(got.Header.Get("Content-Type"))

}
