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
	if len(*d) < 1 {
		return "empty data"
	}

	all := make([]string, len(*d))

	for _, v := range *d {
		s := fmt.Sprintf("%v", v)
		all = append(all, s)
	}

	return strings.Join(all, "\n")
}
