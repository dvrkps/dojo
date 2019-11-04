package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dvrkps/dojo/consul/guard"
	"github.com/dvrkps/dojo/consul/logger"
	"github.com/dvrkps/dojo/consul/randomservice"
)

func main() {
	lgr := logger.New(true, os.Stdout, os.Stderr, 100)
	lgr.Debug("logger set")

	//teardown := initConsulKV(lgr)
	//defer teardown()

	ctx, cancel := context.WithCancel(context.Background())
	stopOnSignal(lgr, cancel)
	lgr.Debug("context and signal set")

	guard := guard.New(ctx)
	for !guard.Quit() {
		startServices(ctx, guard, lgr)
	}
	lgr.Info("the end")
}

func startServices(ctx context.Context, guard *guard.Guard, lgr *logger.Logger) {
	const (
		capRandomResponse    = 100
		randomServiceWorkers = 2
		randomMaxNumber      = 10
	)

	randomResponse := make(chan int, capRandomResponse)
	stopRandomService, err := randomservice.Start(ctx, guard, lgr, randomServiceWorkers, randomMaxNumber, randomResponse)
	if err != nil {
		lgr.Error("random: %v", err)
	}
	defer stopRandomService()

}

func stopOnSignal(lgr *logger.Logger, cancel func()) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-exit
		signal.Stop(exit)
		lgr.Debug("signal")
		cancel()
	}()
}

func initConsulKV(lgr *logger.Logger) func() {
	cc, err := newConsulClient(consulClusterAddress())
	if err != nil {
		lgr.Error("client: %v", err)
		return func() {}
	}

	clearKV, err := consulSetupKV(cc)
	if err != nil {
		lgr.Error("setup: %v", err)
		return func() {}
	}
	return func() {
		err = clearKV()
		if err != nil {
			lgr.Error("clear: %v", err)
		}
		lgr.Debug("teardown end")
	}
}
