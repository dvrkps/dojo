package main

import "fmt"

// Server is custom server.
type Server struct {
	host string
	port int
}

// Host sets server's host.
func Host(host string) func(s *Server) {
	return func(s *Server) {
		s.host = host
	}
}

// Port sets server's port.
func Port(port int) func(s *Server) {
	return func(s *Server) {
		s.port = port
	}
}

func newServer(options ...func(s *Server)) *Server {
	srv := &Server{}
	for _, o := range options {
		o(srv)
	}
	return srv
}

func main() {
	srv := newServer(
		Port(80),
		Host("example.com"),
		Port(82))
	fmt.Printf("%+v\n", srv)
}
