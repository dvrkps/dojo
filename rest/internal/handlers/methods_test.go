package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMethods(t *testing.T) {
	for _, tt := range methodTests() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			tt.h(w, tt.r)

			got := w.Result()
			defer got.Body.Close()

			gotCode := got.StatusCode

			if tt.fail {
				if gotCode != http.StatusMethodNotAllowed {
					t.Errorf("status code = %v; want %v", gotCode, http.StatusMethodNotAllowed)
				}

				return
			}

			if gotCode != http.StatusOK {
				t.Errorf("status code = %v; want %v", gotCode, http.StatusOK)
			}
		})
	}
}

type methodTest struct {
	name string
	fail bool
	r    *http.Request
	h    http.HandlerFunc
}

func methodTests() []methodTest {
	rGet := httptest.NewRequest("GET", "/", nil)
	rPost := httptest.NewRequest("POST", "/", nil)

	dummy := func(w http.ResponseWriter, r *http.Request) {}

	a := API{}
	hGet := a.withGet(dummy)
	hPost := a.withPost(dummy)

	tests := []methodTest{
		{
			name: "ok get",
			r:    rGet,
			h:    hGet,
		},
		{
			name: "ok post",
			r:    rPost,
			h:    hPost,
		},
		{
			name: "no get",
			fail: true,
			r:    rPost,
			h:    hGet,
		},
		{
			name: "no post",
			fail: true,
			r:    rGet,
			h:    hPost,
		},
	}

	return tests
}
