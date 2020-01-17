package handlers

import (
	"fmt"
	"net/http"
)

func (a *API) withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("nutra")
		fmt.Printf("%+v", r)
		//w.WriteHeader(http.StatusUnauthorized)
		next(w, r)
	}
}
