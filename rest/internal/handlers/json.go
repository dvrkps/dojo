package handlers

import "net/http"

const contentTypeHeader = "Content-Type"

const jsonContentType = "application/json; charset=utf-8"

func (a *API) withJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeHeader, jsonContentType)
		next(w, r)
	}
}
