package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"testing"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestRun(t *testing.T) {

	const (
		brokerIP  = "127.0.0.1"
		serviceID = "42"
	)

	tests := []struct {
		s    *Socket
		want int
		bip  string
		sid  string
	}{
		{
			s: &Socket{
				zs: &fakeSocket{},
			},
			want: exitUser,
			bip:  brokerIP,
			sid:  "",
		},
		{
			s: &Socket{
				zs: &fakeSocket{},
			},
			want: exitUser,
			bip:  "",
			sid:  serviceID,
		},
		{
			s: &Socket{
				zs:  &fakeSocket{},
				err: errors.New("error"),
			},
			want: exitErr,
			bip:  brokerIP,
			sid:  serviceID,
		},
		{
			s: &Socket{
				zs: &fakeSocket{
					isConnectErr: true,
				},
			},
			want: exitErr,
			bip:  brokerIP,
			sid:  serviceID,
		},
		{
			s: &Socket{
				zs: &fakeSocket{
					isRecvErr: true,
				},
			},
			want: exitErr,
			bip:  brokerIP,
			sid:  serviceID,
		},
		{
			s: &Socket{
				zs: &fakeSocket{
					recvMsg: "NaN",
				},
			},
			want: exitErr,
			bip:  brokerIP,
			sid:  serviceID,
		},
		{
			s: &Socket{
				zs: &fakeSocket{
					recvMsg: strconv.Itoa(replyLimit),
				},
			},
			want: exitOK,
			bip:  brokerIP,
			sid:  serviceID,
		},
		{
			s: &Socket{
				zs: &fakeSocket{
					recvMsg: strconv.Itoa(replyLimit + 1),
				},
			},
			want: exitOK,
			bip:  brokerIP,
			sid:  serviceID,
		},
		{
			s: &Socket{
				zs: &fakeSocket{
					isSendErr: true,
				},
			},
			want: exitErr,
			bip:  brokerIP,
			sid:  serviceID,
		},
	}

	for _, tt := range tests {

		if got := run(tt.s, tt.bip, tt.sid); got != tt.want {
			t.Errorf("run(...) = %v; want %v", got, tt.want)
		}
	}
}

func TestEndpoint(t *testing.T) {

	port := strconv.Itoa(defaultPort)

	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "1.2.3.4",
			want: "tcp://1.2.3.4:" + port,
		},
	}

	for _, tt := range tests {
		if got, err := Endpoint(tt.in); got != tt.want || err != nil {
			t.Errorf("Endpoint(%q) = %q, %v; want %v, <nil>",
				tt.in, got, err, tt.want,
			)
		}
	}
}

func TestEndpoint_errors(t *testing.T) {

	tests := []struct {
		in string
	}{
		{in: ""},
		{in: "localhost"},
		{in: "fake"},
	}

	for _, tt := range tests {
		if got, err := Endpoint(tt.in); got != "" || err == nil {
			t.Errorf("Endpoint(%q) = %q, %v; want \"\", <error>",
				tt.in, got, err,
			)
		}
	}
}
