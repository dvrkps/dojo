package fileload

import (
	"bytes"
	"fmt"
	"strconv"
)

// Row holds file row's data.
type Row struct {
	ID   int
	Name []byte
}

// NewRow creates new Row.
func NewRow(in []byte) (Row, error) {
	sep := []byte{':'}

	all := bytes.Split(in, sep)

	const fields = 2
	if len(all) < fields {
		return Row{}, fmt.Errorf("fields: len < %v", fields)
	}

	id, err := strconv.Atoi(string(all[0]))
	if err != nil {
		return Row{}, fmt.Errorf("id: %v", err)
	}

	r := Row{
		ID:   id,
		Name: all[1],
	}

	return r, nil
}

// String returns string representation.
func (r *Row) String() string {
	s := fmt.Sprintf("id: %v name: %v", r.ID, r.Name)
	return s
}
