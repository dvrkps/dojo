package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server represents server data.
type Server struct {
	counter int
	quit    chan os.Signal
}

// NewServer init and starts server.
func NewServer() {

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	s := &Server{
		quit: ch,
	}

	s.run()
}

const counterDelay = 1 * time.Second

const maxCount = 42

func (s *Server) run() {

	log.Print("server start")

	for {
		select {

		case <-s.quit:
			s.stop()
			log.Print("server stop, counter:", s.counter)

			return

		default:
			if s.counter > maxCount {
				return
			}
			s.counter++
			log.Print(s.counter)
			time.Sleep(counterDelay)
		}
	}
}

func (s *Server) stop() {

	for s.counter > 0 {
		s.counter--
		log.Print("stop:", s.counter)
		time.Sleep(counterDelay)
	}
}

func main() {
	NewServer()
	fmt.Println("the end")
}
