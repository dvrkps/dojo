package main

type kinder interface {
	kind() string
}

type primeser interface {
	primes() []uint
}

type kindPrimeser interface {
	kinder
	primeser
}

type primeType uint

func (pt primeType) String() string {
	switch pt {
	case baseType:
		return baseTypeName
	case concType:
		return concTypeName
	default:
		return invalidTypeName
	}
}

const (
	baseType primeType = iota + 1
	concType
)

const (
	baseTypeName    = "base"
	concTypeName    = "concurrent"
	invalidTypeName = "invalid"
)
