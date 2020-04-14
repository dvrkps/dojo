package fileload

import "fmt"

// Data holds all file rows.
type Data []Row

// Add new row.
func (d Data) Add(in []byte) (Data, error) {
	r, err := NewRow(in)
	if err != nil {
		return Data{}, fmt.Errorf("%q: %v", in, err)
	}

	d = append(d, r)

	return d, nil
}
