package parallel

import "time"

const delay = 2 * time.Second

// Same returns output same as input.
func Same(i int) int {
	time.Sleep(delay)
	return i
}
