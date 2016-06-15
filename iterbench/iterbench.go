package iterbench

func iterFor(lmt int) int {
	var (
		i   int
		sum int
		ok  bool
	)
	for {
		i, ok = do(i, lmt)
		sum += i
		if !ok {
			break
		}
	}
	return sum
}

func do(i int, lmt int) (int, bool) {
	i++
	ok := i != lmt
	return i, ok
}
