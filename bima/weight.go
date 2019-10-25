package main

const (
	minWeight = 1.0
	maxWeight = 400.0
)

// Weight holds value in kilograms.
type Weight float64

// NewWeight creates weight.
func NewWeight(v float64) (Weight, error) {
	w := Weight(v)
	if err := w.Err(); err != nil {
		return Weight(0), err
	}
	return w, nil
}

// Err returns weight error.
func (w Weight) Err() error {
	return checkValue("weight", float64(w), minWeight, maxWeight)
}
