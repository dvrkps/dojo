package medicament

import (
	"errors"
	"fmt"
	"strconv"
)

type dosage []float64

func (d dosage) ratio() float64 {
	size := float64(len(d))
	if size < 1 {
		return 0
	}

	var sum float64

	for _, v := range d {
		sum += v
	}

	return sum / size
}

func newDosage(doses ...[]byte) (dosage, error) {
	empty := dosage{}

	if len(doses) < 1 {
		return empty, errors.New("dosage: empty")
	}

	all := dosage{}

	for _, d := range doses {
		f, err := strconv.ParseFloat(string(d), 64)

		if err != nil {
			return empty, fmt.Errorf("dosage: %v", err)
		}

		if f < 0 {
			return empty, fmt.Errorf("dosage: %v < 0", f)
		}

		all = append(all, f)
	}

	return all, nil
}
