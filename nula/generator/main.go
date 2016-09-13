package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	var ip = flag.String("ip", "", "broker ip address")
	flag.Parse()

	os.Exit(run(NewSocket(zmq.REQ), *ip, EnvServiceID()))
}

const (
	exitOK   = 0
	exitErr  = 1
	exitUser = 2
)

const replyLimit = 50

func run(skt *Socket, brokerIP, envSid string) int {

	defer CloseSocket(skt)

	sid, err := ServiceID(envSid)
	if err != nil {
		log.Printf("ServiceID: %v", err)
		return exitUser
	}

	brokerEndpoint, err := Endpoint(brokerIP)
	if err != nil {
		log.Printf("endpoint: %v", err)
		return exitUser
	}

	if err := skt.Connect(brokerEndpoint); err != nil {
		log.Printf("connect: %v", err)
		return exitErr
	}

	factory := NewFactory()

	const (
		msgLimit = 30
	)

	for i := 0; i < msgLimit; i++ {

		jobID := factory.NewJobID(time.Now(), sid)

		if _, err := skt.Send(jobID, 0); err != nil {
			log.Printf("send: %v", err)
			return exitErr
		}

		in, err := skt.Recv(0)
		if err != nil {
			log.Printf("recv: %v", err)
			return exitErr
		}

		reply, err := strconv.Atoi(in)
		if err != nil {
			log.Print(err)
			return exitErr
		}

		log.Printf("out: %v in: %d", jobID, reply)

		if reply > replyLimit {
			log.Printf("exit: %v > %v", reply, replyLimit)
			break
		}

	}

	log.Print("The end.")
	return exitOK
}

const defaultPort = 8001

// Endpoint returns valid endpoint address.
func Endpoint(host string) (string, error) {

	ip := net.ParseIP(host)
	if ip == nil {
		return "", fmt.Errorf("ip %q: not valid", host)
	}

	host = fmt.Sprintf("tcp://%v:%d", ip, defaultPort)

	return host, nil
}
