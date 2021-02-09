package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	dojoPings, err := startCounterVec("dojo_pings_total", "Number of pings", []string{"instance"})
	if err != nil {
		log.Printf("start: %v", err)
		return
	}

	go startInstance("one", dojoPings)

	go startInstance("two", dojoPings)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func startCounterVec(name, help string, labels []string) (*prometheus.CounterVec, error) {
	cv := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	err := prometheus.Register(cv)
	if err != nil {
		return nil, err
	}

	return cv, nil
}

func startInstance(name string, c *prometheus.CounterVec) {
	p := c.WithLabelValues(name)
	for {
		p.Inc()
		randomSleep()
	}
}

func randomSleep() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000000000)
	time.Sleep(time.Duration(n))
}
