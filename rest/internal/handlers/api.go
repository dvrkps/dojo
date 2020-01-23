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
	dummy := func(name string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			const randMax = 1000
			n := rand.Intn(randMax)
			fmt.Fprintf(w, "%v: %v\n", name, n)
		}
	}

	m := http.NewServeMux()

	m.HandleFunc("/a", a.withJSON(dummy("A")))
	m.HandleFunc("/b", a.withJSON(dummy("B")))

	a.mux = m

	return a
}

// ServeHTTP implements http.Handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
