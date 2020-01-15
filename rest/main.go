package main

import (
	"io"
	"os"
	"syscall"

	"github.com/dvrkps/dojo/rest/internal/configuration"
	"github.com/dvrkps/dojo/rest/internal/handlers"
	"github.com/dvrkps/dojo/rest/internal/log"
	"github.com/dvrkps/dojo/rest/internal/server"
)

func main() {
	os.Exit(run(os.Args, os.Stdout, os.Stderr))
}

const (
	exitOk   = 0
	exitErr  = 1
	exitUser = 2
)

func run(args []string, stdout, stderr io.Writer) int {
	log := log.New(stdout, stderr)

	cfg, err := configuration.New(args, stderr)
	if err != nil {
		log.Errorf("configuration: %v", err)
		return exitUser
	}

	if cfg.Verbose {
		log.Verbose()
	}

	const (
		apiAddress = "localhost:8000"
	)

	api := handlers.API{}

	s := server.Server{
		Addr:             cfg.Addr,
		Handler:          api.Routes(),
		Log:              log,
		TerminateSignals: []os.Signal{os.Interrupt, syscall.SIGTERM},
	}

	if err := s.Run(); err != nil {
		log.Errorf("server: %v", err)
		return exitErr
	}

	return exitOk
}
