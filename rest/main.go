package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dvrkps/dojo/rest/internal/configuration"
	"github.com/dvrkps/dojo/rest/internal/log"
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

const (
	exitOk   = 0
	exitErr  = 1
	exitUser = 2
)

func run(args []string, stdout, stderr io.Writer) int {
	log := log.New(stdout, stderr)

	cfg, err := configuration.New(args, stderr)
	if err != nil {
		log.Errorf("configuration: %v", err)
		return exitUser
	}
	if cfg.Verbose {
		log.Verbose()
	}

	const (
		apiAddress   = "localhost:8000"
		readTimeout  = 5 * time.Second
		writeTimeout = 5 * time.Second
	)

	api := http.Server{
		Addr:         apiAddress,
		Handler:      newServer(),
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Infof("api listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Errorf("api error: %v", err)
		return exitErr

	case <-shutdown:
		const shutdownTimeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err = api.Shutdown(ctx)
		if err != nil {
			log.Errorf("api shutdown timeout: %v", err)
			err = api.Close()
		}

		if err != nil {
			log.Errorf("api stop: %v", err)
			return exitErr
		}
	}

	return exitOk
}

type server struct {
	router http.Handler
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() *server {
	m := http.NewServeMux()
	m.HandleFunc("/a", jsonContentType(aecho))
	m.HandleFunc("/b", jsonContentType(becho))
	s := server{
		router: m,
	}
	return &s
}

func jsonContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r)
	}
}

func aecho(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(1000)
	fmt.Fprintf(w, "a) You asked to %s %s result: %d\n", r.Method, r.URL.Path, n)
}

func becho(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(1000)
	fmt.Fprintf(w, "b) You asked to %s %s result: %d\n", r.Method, r.URL.Path, n)
}
