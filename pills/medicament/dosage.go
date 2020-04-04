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

const (
	minusChar    = `-`
	quarterValue = 0.25
	quarterChar  = `q`
	halfValue    = 0.5
	halfChar     = `h`
)

// String returns Dosage string representation.
func (d Dosage) String() string {
	result := make([][]byte, len(d.doses))

	for i, n := range d.doses {
		switch {
		case n <= 0:
			result[i] = append(result[i], []byte(minusChar)...)
		case n == quarterValue:
			result[i] = append(result[i], []byte(quarterChar)...)
		case n == halfValue:
			result[i] = append(result[i], []byte(halfChar)...)
		default:
			result[i] = strconv.AppendFloat(result[i], n, 'g', -1, 64)
		}
	}

	j := bytes.Join(result, []byte(" "))

	return string(j)
}

// newDosage creates Dosage.
func newDosage(doses ...[]byte) (Dosage, error) {
	empty := Dosage{}

	if len(doses) < 1 {
		return empty, errors.New("dosage: empty")
	}

	dss := []float64{}

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

	if bytes.EqualFold(in, []byte(halfChar)) {
		return halfValue, nil
	}

	if bytes.EqualFold(in, []byte(quarterChar)) {
		return quarterValue, nil
	}

	result, err := strconv.ParseFloat(string(in), 64)
	if result < 0 || err != nil {
		return 0, doseErr
	}

	return result, nil
}
