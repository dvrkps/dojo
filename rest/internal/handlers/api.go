// Package handlers contains all api handlers.
package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
)

// API holds all routes.
type API struct {
	mux *http.ServeMux
}

// Routes inits all api routes.
func (a *API) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/a", jsonContentType(aecho))
	m.HandleFunc("/b", jsonContentType(becho))

	a.mux = m

	return a
}

// ServeHTTP implements http.Handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func aecho(w http.ResponseWriter, r *http.Request) {
	const randMax = 1000
	n := rand.Intn(randMax)
	fmt.Fprintf(w, "a) You asked to %s %s result: %d\n", r.Method, r.URL.Path, n)
}

func becho(w http.ResponseWriter, r *http.Request) {
	const randMax = 1000
	n := rand.Intn(randMax)
	fmt.Fprintf(w, "b) You asked to %s %s result: %d\n", r.Method, r.URL.Path, n)
}

func jsonContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r)
	}
}
