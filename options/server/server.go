package server

import "fmt"

// Server is custom server.
type Server struct {
	host string
	port int
}

func (s *Server) setHost(host string) error {
	if host == "" {
		return fmt.Errorf("server: %s: invalid host", host)
	}
	s.host = host
	return nil
}

// Host sets server's host.
func Host(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

// Port sets server's port.
func Port(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// Option handle server option.
type Option func(s *Server)

// New creates new server.
func New(options ...Option) *Server {
	srv := &Server{}
	for _, o := range options {
		o(srv)
	}
	return srv
}
