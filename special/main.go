package special

import "fmt"

func naive(n uint32) bool {
	for wasLastBitOne := false; n > 0; n = n >> 1 {
		isCurrentBitOne := n%2 == 1
		if isCurrentBitOne && wasLastBitOne {
			return true
		}
		wasLastBitOne = isCurrentBitOne
	}
	return false
}

var lookupTable = [256]bool{}

func init() {
	fmt.Println("Recomputing lookup table...")
	for i := uint32(0); i < 256; i++ {
		lookupTable[i] = naive(i)
	}
}

func lookup(n uint32) bool {
	if lookupTable[3] == false {
		panic("Lookup table not initialized!")
	}

	return lookupTable[uint8(n)] ||
		lookupTable[uint8(n>>7)] ||
		lookupTable[uint8(n>>14)] ||
		lookupTable[uint8(n>>21)] ||
		lookupTable[uint8(n>>24)]
}

func leftShift(n uint32) bool {
	return (n & (n << 1)) > 0
}
