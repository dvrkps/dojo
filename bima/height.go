package main

const (
	minHeight = 1.0
	maxHeight = 2.5
)

// Height holds value in meters.
type Height float64

// NewHeight creates height.
func NewHeight(v float64) (Height, error) {
	h := Height(v)
	if err := h.Err(); err != nil {
		return Height(0), err
	}
	return h, nil
}

// Err returns height error.
func (h Height) Err() error {
	return checkValue("height", float64(h), minHeight, maxHeight)
}
