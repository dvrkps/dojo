package handlers

import "net/http"

func (a *API) withGet(next http.HandlerFunc) http.HandlerFunc {
	return checkMethod(http.MethodGet, next)
}

func (a *API) withPost(next http.HandlerFunc) http.HandlerFunc {
	return checkMethod(http.MethodPost, next)
}

func checkMethod(methodName string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != methodName {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		next(w, r)
	}
}
