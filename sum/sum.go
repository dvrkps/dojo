package sum

func basic(sum int, nums ...int) bool {
	for _, x := range nums {
		for _, y := range nums {
			v := x + y
			if sum == v {
				return true
			}
		}
	}
	return false
}

func better(sum int, nums ...int) bool {
	for _, x := range nums {
		for _, y := range nums {
			v := x + y
			if sum == v {
				return true
			}
		}
	}
	return false
}
