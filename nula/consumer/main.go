package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	os.Exit(run(NewSocket(zmq.REP)))
}

const (
	exitErr = 1
)

const (
	defaultEndpoint = "tcp://localhost:9001"
)

func run(skt *Socket) int {

	defer CloseSocket(skt)

	if err := skt.Connect(defaultEndpoint); err != nil {
		log.Printf("connect: %v", err)
		return exitErr
	}

	for {
		in, err := skt.Recv(0)
		if err != nil {
			log.Printf("in: %v", err)
			break
		}

		out := Reply()

		log.Printf("in: %s out: %d \n", in, out)

		time.Sleep(waitMs())

		_, err = skt.Send(strconv.Itoa(out), 0)
		if err != nil {
			log.Printf("out: %v", err)
			break
		}
	}

	return exitErr
}

const (
	minReply = 1
	maxReply = 60
)

// Reply returns random number
// between minReply and maxReply.
func Reply() int {
	return randInRange(minReply, maxReply)
}

const (
	minWaitMs = 1
	maxWaitMs = 2000
)

func waitMs() time.Duration {
	r := randInRange(minWaitMs, maxWaitMs)
	return time.Duration(r) * time.Millisecond
}

func randInRange(min, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(max-min) + min
}
