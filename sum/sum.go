package sum

func basic(sum int, nums ...int) bool {
	for _, x := range nums {
		for _, y := range nums {
			s := x + y
			if sum == s {
				return true
			}
		}
	}
	return false
}

func better(sum int, nums ...int) bool {
	max := len(nums)
	for ix := 0; ix < max; ix++ {
		x := nums[ix]
		for iy := ix + 1; iy < max; iy++ {
			y := nums[iy]
			s := x + y
			if sum == s {
				return true
			}
		}
	}
	return false
}

func linear(sum int, nums ...int) bool {
	min := 0
	max := len(nums) - 1
	for min < max {
		s := nums[min] + nums[max]
		switch {
		case s == sum:
			return true
		case s > sum:
			max--
		case s < sum:
			min++
		}
	}
	return false
}
