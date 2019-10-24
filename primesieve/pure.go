package main

func pure(max uint) []uint {
	exists := make(map[uint]struct{}, max)
	primes := make([]uint, 0, max)
	for i := uint(2); i < max; i++ {
		if _, ok := exists[i]; ok {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < max; k += i {
			exists[k] = struct{}{}
		}
	}
	return primes
}
