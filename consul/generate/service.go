package generate

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/dvrkps/consuldojo/logger"
)

// Service is generator service.
type Service struct {
	NoWorkers  int
	Maximum    int
	Log        *logger.Logger
	DoneChan   <-chan struct{}
	ResultChan chan Result
}

func (s *Service) validate() error {
	if s.NoWorkers < 0 {
		return errors.New("NoWorkers < 0")
	}
	if s.Maximum < 1 {
		return errors.New("Maximum < 1")
	}
	if s.DoneChan == nil {
		return errors.New("nil done channel")
	}
	if s.Log == nil {
		return errors.New("nil logger")
	}
	if s.ResultChan == nil {
		return errors.New("nil result channel")
	}
	return nil
}

// Start runs service.
func (s *Service) Start() error {
	err := s.validate()
	if err != nil {
		return fmt.Errorf("generate: start failed: %v", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(s.NoWorkers)
	for i := 0; i < s.NoWorkers; i++ {
		go s.runWorker(&wg, i)
	}
	s.Log.Info("generate: start %v workers", s.NoWorkers)
	wg.Wait()
	close(s.ResultChan)
	s.Log.Info("generate: stop all workers")
	return nil
}

func (s *Service) runWorker(wg *sync.WaitGroup, id int) {
	s.Log.Debug("generate: run %v worker", id)
	defer wg.Done()
	rnd := newRand()
	for {
		select {
		case <-s.DoneChan:
			s.Log.Debug("generate: stop %v worker", id)
			return
		default:
			r, err := newResult(int64(id), int64(s.Maximum), rnd)
			if err != nil {
				s.Log.Debug("generate: worker %v: %v", id, err)
				continue
			}
			s.ResultChan <- r
		}
	}
}

func newRand() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s)
}
