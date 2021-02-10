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
	dojoRestarts, err := startCounterVec("dojo_restarts_total", "Number of restarts", []string{"service"})
	if err != nil {
		log.Printf("start: %v", err)
		return
	}

	go startService("one", dojoRestarts)

	go startService("two", dojoRestarts)

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

func startService(name string, restarts *prometheus.CounterVec) {
	for {
		done := make(chan struct{})
		t := time.NewTimer(1 * time.Minute)

		go func() {
			<-t.C
			close(done)
		}()

		const maxSleep = 1000000000
		restarts.WithLabelValues(name).Inc()
		for {
			rs := randomNumber(maxSleep)
			time.Sleep(time.Duration(rs))
		}

		println("restart " + name)
	}
}

func randomNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
