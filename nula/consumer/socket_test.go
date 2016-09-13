package main

import (
	"errors"
	"testing"

	zmq "github.com/pebbe/zmq4"
)

func TestNewSocket(t *testing.T) {
	const reqType = zmq.REQ
	if s := NewSocket(reqType); s.err != nil {
		t.Errorf("NewSocket(%v) error %v; want <nil>",
			reqType, s.err)
	}

	const invalidType = -1
	if s := NewSocket(invalidType); s.err == nil {
		t.Errorf("NewSocket(%v) error %v; want <error>",
			reqType, s.err)
	}
}

type fakeSocket struct {
	isConnectErr bool
	isCloseErr   bool
	isRecvErr    bool
	recvMsg      string
	isSendErr    bool
	sendInt      int
}

func (s *fakeSocket) Connect(endpoint string) error {
	if s.isConnectErr {
		return errors.New("connect error")
	}
	return nil
}

func (s *fakeSocket) Close() error {
	if s.isCloseErr {
		return errors.New("close error")
	}
	return nil
}

func (s *fakeSocket) Recv(flags zmq.Flag) (string, error) {
	if s.isRecvErr {
		return s.recvMsg, errors.New("recv error")
	}
	return s.recvMsg, nil
}

func (s *fakeSocket) Send(data string, flags zmq.Flag) (int, error) {
	if s.isSendErr {
		return 0, errors.New("recv error")
	}
	return s.sendInt, nil
}

// Raise error on compile if fakeSocket
// not implements Socket.
var _ *Socket = &Socket{zs: &fakeSocket{}}

func TestSocket_Err(t *testing.T) {
	s := &Socket{
		err: errors.New("fake error"),
	}
	if s.Err() == nil {
		t.Errorf("(%v) Err() = <nil>; want <error>",
			s.err,
		)
	}
}

func TestSocket_Connect(t *testing.T) {
	tests := []*Socket{
		{
			zs: &fakeSocket{},
		},
	}

	for _, tt := range tests {

		if got := tt.Connect("endpoint"); got != nil {
			t.Errorf("Connect(...) = %v; want <nil>", got)
		}
	}
}

func TestSocket_Connect_errors(t *testing.T) {
	tests := []*Socket{
		{
			zs:  &fakeSocket{},
			err: errors.New("error"),
		},
		{
			zs: &fakeSocket{
				isConnectErr: true,
			},
		},
	}

	for _, tt := range tests {

		if got := tt.Connect("endpoint"); got == nil {
			t.Errorf("Connect(...) = %v; want <error>", got)
		}
	}
}

func TestSocket_Close(t *testing.T) {
	tests := []*Socket{
		{
			zs: &fakeSocket{},
		},
	}

	for _, tt := range tests {

		if got := tt.Close(); got != nil {
			t.Errorf("Close() = %v; want <nil>", got)
		}
	}
}

func TestSocket_Close_errors(t *testing.T) {
	tests := []*Socket{
		{
			zs:  &fakeSocket{},
			err: errors.New("error"),
		},
		{
			zs: &fakeSocket{
				isCloseErr: true,
			},
		},
	}

	for _, tt := range tests {

		if got := tt.Close(); got == nil {
			t.Errorf("Close() = %v; want <error>", got)
		}
	}
}

func TestSocket_Recv(t *testing.T) {

	tests := []struct {
		s    *Socket
		want string
	}{
		{
			s: &Socket{
				zs: &fakeSocket{
					recvMsg: "abra",
				},
			},
			want: "abra",
		},
	}

	for _, tt := range tests {

		if got, err := tt.s.Recv(0); got != tt.want || err != nil {
			t.Errorf("Recv(0) = %v, %v; want %v, <nil>",
				got, err, tt.want)
		}
	}
}

func TestSocket_Recv_errors(t *testing.T) {
	tests := []*Socket{
		{
			zs:  &fakeSocket{},
			err: errors.New("error"),
		},
		{
			zs: &fakeSocket{
				isRecvErr: true,
			},
		},
	}

	for _, tt := range tests {

		if _, err := tt.Recv(0); err == nil {
			t.Errorf("Recv(0) = _, %v; want _, <error>",
				err)
		}
	}
}

func TestSocket_Send(t *testing.T) {

	tests := []struct {
		s    *Socket
		want int
	}{
		{
			s: &Socket{
				zs: &fakeSocket{
					sendInt: 42,
				},
			},
			want: 42,
		},
	}

	for _, tt := range tests {

		if got, err := tt.s.Send("message", 0); got != tt.want || err != nil {
			t.Errorf("Send(\"message\", 0) = %v, %v; want %v, <nil>",
				got, err, tt.want)
		}
	}
}

func TestSocket_Send_errors(t *testing.T) {
	tests := []*Socket{
		{
			zs:  &fakeSocket{},
			err: errors.New("error"),
		},
		{
			zs: &fakeSocket{
				isSendErr: true,
			},
		},
	}

	for _, tt := range tests {

		if _, err := tt.Send("", 0); err == nil {
			t.Errorf("Send(\"\", 0) = _, %v; want _, <error>",
				err)
		}
	}
}

func TestCloseSocket(t *testing.T) {

	tests := []*Socket{
		nil,
		{
			zs:  &fakeSocket{},
			err: errors.New("error"),
		},
		{
			zs: &fakeSocket{
				isCloseErr: true,
			},
		},
	}

	for _, s := range tests {
		CloseSocket(s)
	}
}
