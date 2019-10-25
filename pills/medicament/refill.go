package medicament

import (
	"errors"
	"strconv"
	"time"
)

type refill struct {
	Date     time.Time
	Quantity float64
}

func newRefill(date, quantity []byte) (refill, error) {
	empty := refill{}
	// date
	d, err := time.Parse("2006-01-02", string(date))
	if err != nil {
		return empty, errors.New("refill: invalid date")
	}
	// quantity
	q, err := strconv.ParseFloat(string(quantity), 64)
	if err != nil {
		return empty, errors.New("refill: invalid quantity")
	}
	if q < 0 {
		return empty, errors.New("refill: quantity less than zero")
	}
	// result
	s := refill{
		Date:     midnight(d),
		Quantity: q,
	}
	return s, nil
}
