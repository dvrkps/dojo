package main

import (
	"errors"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

func TestRun(t *testing.T) {

	tests := []struct {
		s    *Socket
		want int
	}{
		{
			s: &Socket{
				zs:  &fakeSocket{},
				err: errors.New("error"),
			},
			want: exitErr,
		},

		{
			s: &Socket{
				zs: &fakeSocket{
					isConnectErr: true,
				},
			},
			want: exitErr,
		},
		{
			s: &Socket{
				zs: &fakeSocket{
					isRecvErr: true,
				},
			},
			want: exitErr,
		},
		{
			s: &Socket{
				zs: &fakeSocket{
					isSendErr: true,
				},
			},
			want: exitErr,
		},
	}

	for _, tt := range tests {

		if got := run(tt.s); got != tt.want {
			t.Errorf("run(...) = %v; want %v", got, tt.want)
		}
	}
}

func TestWaitMs(t *testing.T) {

	min := minWaitMs * time.Millisecond
	max := maxWaitMs * time.Millisecond

	for i := 0; i < 1000; i++ {
		got := waitMs()
		if got < min || got > max {
			t.Fatalf("waitMs() = %v; want %v < x < %v",
				got, min, max)
		}
	}
}

func TestReply(t *testing.T) {

	for i := 0; i < 1000; i++ {
		got := Reply()
		if got < minReply || got > maxReply {
			t.Fatalf("Reply() = %v; want %v < x < %v",
				got, minReply, maxReply)
		}
	}
}
