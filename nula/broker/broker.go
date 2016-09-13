package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

// Broker connects dealer and router.
type Broker struct {
	D *zmq.Socket
	R *zmq.Socket
	p *zmq.Poller
}

// NewBroker creates broker.
func NewBroker() (*Broker, error) {

	rs, err := zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		return nil, fmt.Errorf("router: %v", err)
	}

	ds, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		return nil, fmt.Errorf("dealer: %v", err)
	}

	p := zmq.NewPoller()
	p.Add(rs, zmq.POLLIN)
	p.Add(ds, zmq.POLLIN)

	b := &Broker{
		D: ds,
		R: rs,
		p: p,
	}

	return b, nil
}

// Close closes all sockets.
func (b *Broker) Close() {
	if err := b.R.Close(); err != nil {
		log.Printf("router close: %v", err)
	}
	if err := b.D.Close(); err != nil {
		log.Printf("dealer close: %v", err)
	}
}

// Sockets pulls all sockets from puller.
func (b *Broker) Sockets() ([]zmq.Polled, error) {
	return b.p.Poll(-1)
}

// Transmit transmits message between sockets.
func (b *Broker) Transmit(in *zmq.Socket) error {
	typ, err := in.GetType()
	if err != nil {
		return err
	}
	out := b.D
	if typ == zmq.DEALER {
		out = b.R
	}
	// TODO(dvrkps): clean this mess.
	for {
		msg, err := in.Recv(0)
		if err != nil {
			return err
		}
		if more, err := in.GetRcvmore(); more {
			if err != nil {
				return err
			}
			_, err := out.Send(msg, zmq.SNDMORE)
			if err != nil {
				return err
			}
		} else {
			_, err := out.Send(msg, 0)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}
