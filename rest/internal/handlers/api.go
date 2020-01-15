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
	m.HandleFunc("/a", aecho)
	m.HandleFunc("/b", becho)
	a.mux = m

	return a
}

// ServeHTTP implements http.Handler.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func aecho(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(1000)
	fmt.Fprintf(w, "a) You asked to %s %s result: %d\n", r.Method, r.URL.Path, n)
}

func becho(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(1000)
	fmt.Fprintf(w, "b) You asked to %s %s result: %d\n", r.Method, r.URL.Path, n)
}
