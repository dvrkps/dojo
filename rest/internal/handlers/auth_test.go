package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	for _, tt := range authTests() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/", nil)
			r.Header = tt.header

			w := httptest.NewRecorder()

			dummy := func(w http.ResponseWriter, r *http.Request) {}

			a := API{}
			handler := a.withAuth(dummy)
			handler(w, r)

			got := w.Result()
			defer got.Body.Close()

			want := tt.statusCode
			if got.StatusCode != want {
				t.Errorf("withAuth status code = %v; want %v", got.StatusCode, want)
			}
		})
	}
}

type authTest struct {
	name       string
	statusCode int
	header     http.Header
}

const testSignatureValue = "druUFhOEGiAir3wspPxJF3mRU5CaXP0J7LIrN2tVDV8="

func authTests() []authTest {
	tests := []authTest{
		{
			name:       "valid",
			statusCode: http.StatusOK,
			header: http.Header{
				"Date": []string{"Tue, 07 Jun 2011 20:51:35 GMT"},
				"Authorization": []string{
					`algorithm="hmac-sha256",` +
						`headers="date",` +
						`signature="` +
						testSignatureValue +
						`",` +
						`apikey="_here_is_the_api_key_"`,
				},
			},
		},
		{
			name:       "no authorization header",
			statusCode: http.StatusUnauthorized,
			header: http.Header{
				"Date": []string{"Tue, 07 Jun 2011 20:51:35 GMT"},
			},
		},
	}

	return tests
}
