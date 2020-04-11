package fileload

import (
	"fmt"
	"strings"
)

// Data holds all file rows.
type Data []Row

// NewData creates new data.
func NewData() Data {
	return Data{}
}

func (d *Data) Add(in []byte) error {
	r, err := NewRow(in)
	if err != nil {
		return fmt.Errorf("%q: %v", in, err)
	}

	*d = append(*d, r)

	return nil
}

// String returns string representation.
func (d *Data) String() string {
	size := len(*d)
	if size < 1 {
		return "empty data"
	}

	all := make([]string, 0, size)

	for _, r := range *d {
		s := fmt.Sprintf("%v", &r)
		all = append(all, s)
	}

	return strings.Join(all, "\n")
}
