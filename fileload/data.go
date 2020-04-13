package fileload

import (
	"fmt"
	"strings"
)

// String returns string representation.
func (d Data) String() string {
	size := len(d)
	if size < 1 {
		return "empty data"
	}

	all := make([]string, 0, size)

	for _, r := range d {
		s := fmt.Sprintf("%s", r)
		all = append(all, s)
	}

	return strings.Join(all, "\n")
}
