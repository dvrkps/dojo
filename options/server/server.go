package server

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
