package server

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/dvrkps/dojo/rest/internal/log"
)

func TestRun(t *testing.T) {
	for _, tt := range runTests() {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			srvErr := make(chan error, 1)
			go func() {
				srvErr <- tt.s.Run()
			}()
			go func() {
				if !tt.fail {
					const timeout = 100 * time.Microsecond
					time.Sleep(timeout)
					interruptServer(t)
				}
			}()

			err := <-srvErr
			if tt.fail {
				if err == nil {
					t.Error("fail: got nil; want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("error: got %v; want nil", err)
			}
		})
	}
}

func interruptServer(t *testing.T) {
	pid := os.Getpid()

	p, err := os.FindProcess(pid)
	if err != nil {
		t.Error("find", err)
	}

	err = p.Signal(os.Interrupt)
	if err != nil {
		t.Log("signal", err)
	}
}

type runTest struct {
	name string
	fail bool
	s    Server
}

func runTests() []runTest {
	const testServerAddr = "localhost:8080"

	dummy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	tests := []runTest{
		{
			name: "valid",
			s: Server{
				Addr:    testServerAddr,
				Handler: dummy,
				Log:     &log.Log{},
			},
		},
		{
			name: "empty addr",
			fail: true,
			s: Server{
				Addr:    "",
				Handler: dummy,
				Log:     &log.Log{},
			},
		},
		{
			name: "nil handler",
			fail: true,
			s: Server{
				Addr: testServerAddr,
				Log:  &log.Log{},
			},
		},
		{
			name: "nil log",
			fail: true,
			s: Server{
				Addr:    testServerAddr,
				Handler: dummy,
			},
		},
	}

	return tests
}
