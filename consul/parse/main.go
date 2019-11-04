package parse

import "github.com/dvrkps/dojo/consul/logger"

type generateOutput struct {
	Worker    int64 `json:"worker,string"`
	Timestamp int64 `json:"timestamp,string"`
	Number    int64 `json:"number,string"`
}

// Service is parse service.
type Service struct {
	NoWorkers  int
	Log        *logger.Logger
	ResultChan chan Result
}
