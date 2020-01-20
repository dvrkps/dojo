package handlers

import (
	"net/http"
)

func (a *API) withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("%+v", r.Header)
		//w.WriteHeader(http.StatusUnauthorized)
		//return
		next(w, r)
	}
}
