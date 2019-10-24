package guard

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/dvrkps/consuldojo/logger"
)

// Guard guards quit.
type Guard struct {
	sig      chan os.Signal
	done     bool
	quitFunc func()
}

// New creates guard.
func New(log *logger.Logger) (*Guard, error) {
	if log == nil {
		return nil, errors.New("nil logger")
	}
	g := Guard{
		sig: make(chan os.Signal),
	}
	signal.Notify(g.sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-g.sig
		log.Debug("signal")
		signal.Stop(g.sig)
		g.done = true
		if g.quitFunc != nil {
			g.quitFunc()
			log.Debug("quit function invoke")
		}
	}()
	return &g, nil
}

// Reload reports wheather the quit signal is send.
func (g *Guard) Reload() bool {
	return !g.done
}

// OnQuit will invoke function on quit.
func (g *Guard) OnQuit(fn func()) {
	g.quitFunc = fn
}

// Quit invoke quit.
func (g *Guard) Quit() {
	close(g.sig)
}
