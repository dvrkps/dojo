package main

import "fmt"

// Server is custom server.
type Server struct {
	domain string
	port   int
}

// Domain sets server's domain.
func Domain(domain string) func(s *Server) {
	return func(s *Server) {
		s.domain = domain
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
		Domain("example.com"),
		Port(82))
	fmt.Printf("%+v\n", srv)
}
