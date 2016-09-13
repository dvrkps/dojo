package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	os.Exit(run())
}

const (
	exitOK   = 0
	exitErr  = 1
	exitUser = 2
)

const routerEndpoint = "tcp://*:8001"

func run() int {
	br, err := NewBroker()
	if err != nil {
		return exitErr
	}

	defer br.Close()

	if err := br.R.Bind(routerEndpoint); err != nil {
		log.Printf("router bind: %v", err)
		return exitErr
	}

	consumers, err := getConsumers("consumers.txt")
	if err != nil {
		log.Printf("consumers: %v", err)
		return exitUser
	}

	for _, ep := range consumers {
		if err := br.D.Bind(ep); err != nil {
			log.Printf("dealer bind: %v", err)
			return exitErr
		}
	}

	for {
		sockets, err := br.Sockets()
		if err != nil {
			log.Printf("poller.Poll: %v", err)
			break
		}
		for _, s := range sockets {
			if err := br.Transmit(s.Socket); err != nil {
				log.Printf("trans: %v", err)
				return exitErr
			}
		}
	}

	return exitErr
}

func getConsumers(path string) ([]string, error) {
	all := []string{}

	c, err := ioutil.ReadFile(path)
	if err != nil {
		return all, err
	}

	s := bufio.NewScanner(bytes.NewReader(c))

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		ep, err := Endpoint(line)
		if err != nil {
			continue
		}
		all = append(all, ep)
	}
	if s.Err() != nil {
		return all, s.Err()
	}
	return all, nil

}

const consumersPort = 9001

// Endpoint returns valid endpoint address.
func Endpoint(host string) (string, error) {

	ip := net.ParseIP(host)
	if ip == nil {
		return "", fmt.Errorf("ip %q: not valid", host)
	}

	host = fmt.Sprintf("tcp://%v:%d", ip, consumersPort)

	return host, nil
}
