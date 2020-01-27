// Package server wraps http server.
package server

import (
	"context"
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
	}

	return nil
}
