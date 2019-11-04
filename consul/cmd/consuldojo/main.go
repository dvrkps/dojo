package main

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/dvrkps/dojo/consul/guard"
	"github.com/dvrkps/dojo/consul/logger"
)

func main() {
	log := logger.New(true, os.Stdout, os.Stderr)
	g, err := guard.New(log)
	if err != nil {
		log.Error("guard: %v", err)
	}
	for g.Reload() {
		println("\n\nstart loop")
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(5e9)
			log.Info("\n\nreload\n\n")
			cancel()
		}()
		g.OnQuit(func() {
			cancel()
			log.Debug("on quit")
		})
		run(ctx, log)
	}
	log.Info("the end")
}

type starter interface {
	Start() error
}

func run(ctx context.Context, log *logger.Logger) {
	wg := sync.WaitGroup{}
	const workers = 2
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(ctx context.Context, id int, log *logger.Logger) {
			log.Info("run %v worker", id)
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Info("worker %v close", id)
					return
				default:
				}
			}
		}(ctx, i, log)
	}
	wg.Wait()
	log.Info("all service closed")
}
