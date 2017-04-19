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
	max := len(nums)
	var x, y, s int
	for ix := 0; ix < max; ix++ {
		x = nums[ix]
		for iy := ix + 1; iy < max; iy++ {
			y = nums[iy]
			s = x + y
			if sum == s {
				return true
			}
		}
	}
	return false
}
