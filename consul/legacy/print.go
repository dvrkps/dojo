package main

import (
	"errors"

	"github.com/dvrkps/consuldojo/logger"
)

type Printer interface {
	Print() string
}

type printService struct {
	lgr *logger.Logger
}

func newPrintService(lgr *logger.Logger, noWorkers int, request <-chan Printer) (*printService, error) {
	if lgr == nil {
		return nil, errors.New("nil logger")
	}
	if noWorkers < 1 {
		return nil, errors.New("noWorkers < 1")
	}
	if request == nil {
		return nil, errors.New("nil request chan")
	}

	svc := printService{lgr: lgr}

	return &svc, nil

}
