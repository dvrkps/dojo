package special

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
