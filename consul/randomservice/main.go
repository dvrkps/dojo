package randomservice

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/dvrkps/consuldojo/guard"
	"github.com/dvrkps/consuldojo/logger"
)

const defaultMaxNumber = 100

func Start(ctx context.Context, guard *guard.Guard, lgr *logger.Logger, noWorkers int, maxNumber int, out chan<- int) (func(), error) {
	emptyFunc := func() {}
	if ctx == nil {
		return emptyFunc, errors.New("nil context")
	}
	if guard == nil {
		return emptyFunc, errors.New("nil guard")
	}
	if lgr == nil {
		return emptyFunc, errors.New("nil logger")
	}
	if noWorkers < 1 {
		return emptyFunc, errors.New("workers < 1")
	}
	if maxNumber < 1 {
		maxNumber = defaultMaxNumber
	}
	if out == nil {
		return emptyFunc, errors.New("nil out")
	}

	guard.Add()
	var wg sync.WaitGroup
	wg.Add(noWorkers)

	stop := func() {
		lgr.Debug("random: wait close")
		wg.Wait()
		close(out)
		lgr.Debug("random: close out")
		guard.Done()
	}

	for i := 0; i < noWorkers; i++ {
		go worker(ctx, &wg, lgr, i, maxNumber, out)
	}
	return stop, nil
}

func worker(ctx context.Context, wg *sync.WaitGroup, lgr *logger.Logger, id int, maxNumber int, out chan<- int) {
	lgr.Info("random #%d: init worker", id)
	defer func() {
		wg.Done()
		lgr.Debug("random #%d: close worker", id)
	}()

	src := rand.NewSource(time.Now().UnixNano())
	for {
		select {
		case <-ctx.Done():
			lgr.Debug("random #%d: signal", id)
			return
		default:
		}
		r := rand.New(src)
		nbr := r.Intn(maxNumber)
		out <- nbr
		//lgr.Debug("random #%d: number %d", id, nbr)
		time.Sleep(1e9)
	}
}
