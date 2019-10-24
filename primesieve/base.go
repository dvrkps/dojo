package main

type baseGenerator struct {
	max uint
	pt  primeType
}

func newBase(max uint) *baseGenerator {
	return &baseGenerator{max: max, pt: baseType}
}

func (b *baseGenerator) primes() []uint {
	exists := make(map[uint]struct{}, b.max)
	primes := make([]uint, 0, b.max)
	for i := uint(2); i < b.max; i++ {
		if _, ok := exists[i]; ok {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < b.max; k += i {
			exists[k] = struct{}{}
		}
	}
	return primes
}

func (b *baseGenerator) kind() string {
	return b.pt.String()
}
