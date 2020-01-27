// Package server wraps http server.
package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dvrkps/dojo/rest/internal/log"
)

// Server is http server with graceful shutdown.
type Server struct {
	Addr    string
	Handler http.Handler
	Log     *log.Log
}

// Run runs the server.
func (s *Server) Run() error {
	err := s.checkOptions()
	if err != nil {
		return err
	}

	hs := s.newHTTPServer()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	srvErr := make(chan error, 1)

	go func() {
		s.Log.Infof("listening on %s", s.Addr)
		srvErr <- hs.ListenAndServe()
	}()

	select {
	case err := <-srvErr:
		return fmt.Errorf("http: %v", err)
	case <-shutdown:
		return s.shutdown(hs)
	}
}

func (s *Server) checkOptions() error {
	if s.Addr == "" {
		return errors.New("empty addr")
	}

	if s.Handler == nil {
		return errors.New("nil handler")
	}

	if s.Log == nil {
		return errors.New("nil log")
	}

	return nil
}

func (s *Server) newHTTPServer() *http.Server {
	const (
		readTimeout  = 5 * time.Second
		writeTimeout = 5 * time.Second
	)

	hs := http.Server{
		Addr:         s.Addr,
		Handler:      s.Handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return &hs
}

func (s *Server) shutdown(hs *http.Server) error {
	const shutdownTimeout = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err := hs.Shutdown(ctx)
	if err != nil {
		s.Log.Errorf("shutdown: %v", err)
		err = hs.Close()
	}

	if err != nil {
		return fmt.Errorf("close: %v", err)
	}

	return nil
}
