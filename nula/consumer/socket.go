package main

import (
	"log"

	zmq "github.com/pebbe/zmq4"
)

// Socket wraps zmq.Socket.
type Socket struct {
	zs interface {
		Connect(endpoint string) error
		Close() error
		Recv(flags zmq.Flag) (string, error)
		Send(data string, flags zmq.Flag) (int, error)
	}
	err error
}

// NewSocket creates Socket.
func NewSocket(typ zmq.Type) *Socket {
	s, err := zmq.NewSocket(typ)
	return &Socket{zs: s, err: err}
}

// Err returns socket error.
func (s *Socket) Err() error {
	return s.err
}

// Connect creates outgoint connection.
func (s *Socket) Connect(endpoint string) error {
	if s.err != nil {
		return s.err
	}
	return s.zs.Connect(endpoint)
}

// Close closing socket.
func (s *Socket) Close() error {
	if s.err != nil {
		return s.err
	}
	return s.zs.Close()
}

// Recv receive a message part from a socket.
func (s *Socket) Recv(flags zmq.Flag) (string, error) {
	if s.err != nil {
		return "", s.err
	}
	return s.zs.Recv(flags)
}

// Send a message part on a socket.
func (s *Socket) Send(data string, flags zmq.Flag) (int, error) {
	if s.err != nil {
		return 0, s.err
	}
	return s.zs.Send(data, flags)
}

// CloseSocket closing socket.
func CloseSocket(s *Socket) {

	if s == nil {
		log.Printf("nil socket")
		return
	}

	if err := s.Close(); err != nil {
		s.err = err
		log.Printf("socket close: %v", err)
	}
}
