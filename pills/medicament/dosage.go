package medicament

import (
	"bytes"
	"errors"
	"strconv"
)

// calcDosageRatio calculate ratio from sum of doses.
func calcDosageRatio(in ...float64) float64 {
	size := len(in)
	if size < 1 {
		return 0
	}
	var sum float64
	for _, v := range in {
		sum += v
	}
	if sum < 0 {
		return 0
	}
	return sum / float64(size)
}

// Dosage holds medicament dosage.
type Dosage struct {
	doses []float64
	ratio float64
}

// Ratio returns calculated doses ratio.
func (d Dosage) Ratio() float64 {
	return d.ratio
}

// String returns Dosage string representation.
func (d Dosage) String() string {
	const (
		minus   = `-`
		quarter = `q`
		half    = `h`
	)
	var result, sep []byte
	for _, n := range d.doses {
		result = append(result, sep...)
		switch {
		case n <= 0:
			result = append(result, minus...)
		case n == 0.25:
			result = append(result, quarter...)
		case n == 0.5:
			result = append(result, half...)
		default:
			result = strconv.AppendFloat(result, n, 'g', -1, 64)
		}
		// separator char
		if sep == nil {
			sep = []byte{' '}
		}
	}
	return string(result)
}

// newDosage creates Dosage.
func newDosage(doses ...[]byte) (Dosage, error) {
	empty := Dosage{}
	var dss []float64
	if len(doses) < 1 {
		return empty, errors.New("dosage: empty")
	}
	for _, d := range doses {
		nd, err := newDose(d)
		if err != nil {
			return empty, err
		}
		dss = append(dss, nd)
	}
	return Dosage{
		doses: dss,
		ratio: calcDosageRatio(dss...),
	}, nil
}

// newDose create daily dose.
func newDose(in []byte) (float64, error) {
	dot := []byte{'.'}
	doseErr := errors.New("dosage: invalid dose")
	if bytes.HasPrefix(in, dot) || bytes.HasSuffix(in, dot) {
		return 0, doseErr
	}
	half := []byte{'h'}
	quarter := []byte{'q'}
	if bytes.EqualFold(in, half) {
		return 0.5, nil
	}
	if bytes.EqualFold(in, quarter) {
		return 0.25, nil
	}
	result, err := strconv.ParseFloat(string(in), 64)
	if result < 0 || err != nil {
		return 0, doseErr
	}
	return result, nil
}
