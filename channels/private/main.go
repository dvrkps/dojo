package main

import (
	"sync"
)

func main() {
	gs := genService{}
	ps := procService{}
	go ps.Run(3, gs.out)
	gs.Run(6)
}

type genService struct {
	ch chan int
}

const max = 10

func (g *genService) Run(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	g.ch = make(chan int)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			for x := 0; x < max; x++ {
				g.ch <- (i * 1000) + x
			}
		}(i)
	}
	wg.Wait()
	close(g.ch)
}

func (gs *genService) out() (int, bool) {
	ch, ok := <-gs.ch
	return ch, ok
}

type procService struct{}

func (ps *procService) Run(n int, fn func() (int, bool)) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				ch, ok := fn()
				if !ok {
					println(i, ": end")
					return
				}
				println(i, ": ", ch)
			}
		}(i)
	}
	wg.Wait()
}
