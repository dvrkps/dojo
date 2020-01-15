// Package server wraps http server.
package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server is http server with graceful shutdown.
type Server struct {
	Addr             string
	TerminateSignals []os.Signal
}

// Run runs the server.
func (s *Server) Run() error {
	const (
		readTimeout  = 5 * time.Second
		writeTimeout = 5 * time.Second
	)

	hs := http.Server{
		Addr:         s.Addr,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, s.TerminateSignals...)

	srvErr := make(chan error, 1)

	go func() {
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
			err = hs.Close()
		}

		if err != nil {
			return fmt.Errorf("close: %v", err)
		}
	}

	return nil
}
