package handlers

import (
	"net/http"

	"github.com/dvrkps/dojo/rest/internal/authorization"
)

func (a *API) withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authorization.Validate(r.Header)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		next(w, r)
	}
}
