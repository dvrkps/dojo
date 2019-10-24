package main

type concurrentGenerator struct {
	max uint
	pt  primeType
}

func newConcurrent(max uint) *concurrentGenerator {
	return &concurrentGenerator{max: max, pt: concType}
}

func (c *concurrentGenerator) kind() string {
	return c.pt.String()
}

func (c *concurrentGenerator) generate() <-chan uint {
	ch := make(chan uint)

	go func() {
		for i := uint(2); i < c.max; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func (c *concurrentGenerator) filter(in <-chan uint, prime uint) chan uint {
	out := make(chan uint)
	go c.filterGoroutine(in, out, prime)
	return out
}

func (c *concurrentGenerator) filterGoroutine(in <-chan uint, out chan<- uint, prime uint) {
	for {
		i, ok := <-in
		if !ok {
			break
		}
		if i%prime != 0 {
			out <- i
		}
	}
	close(out)
}

func (c *concurrentGenerator) primes() []uint {
	//func concurrent(max int) []int {
	primes := []uint{}
	in := c.generate()
	for {
		p, ok := <-in
		if !ok {
			break
		}
		primes = append(primes, p)
		in = c.filter(in, p)
	}

	return primes
}
