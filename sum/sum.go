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

func complements(sum int, nums ...int) bool {
	comp := map[int]struct{}{}
	for _, n := range nums {
		v := sum - n
		if _, ok := comp[v]; ok {
			return true
		}
		comp[v] = struct{}{}
	}
	return false
}

func optim(sum int, nums ...int) bool {
	comp := make([]int, 0, len(nums))
	for _, n := range nums {
		v := sum - n
		for _, c := range comp {
			if c == v {
				return true
			}
		}
		comp = append(comp, v)
	}
	return false
}
