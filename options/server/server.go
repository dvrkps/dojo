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

func (s *Server) setPort(port int) error {
	if port < 0 {
		return fmt.Errorf("server: invalid port %d", port)
	}
	s.port = port
	return nil
}

// Option handle server option.
type Option func(s *Server) error

// Host sets server's host.
func Host(host string) Option {
	return func(s *Server) error {
		return s.setHost(host)
	}
}

// Port sets server's port.
func Port(port int) Option {
	return func(s *Server) error {
		return s.setPort(port)
	}
}

// New creates new server.
func New(options ...Option) (*Server, error) {
	srv := &Server{}
	for _, o := range options {
		if err := o(srv); err != nil {
			return nil, err
		}
	}
	return srv, nil
}
