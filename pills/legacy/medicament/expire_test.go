package medicament

import (
	"reflect"
	"testing"
)

func fakeExpire(y, m, d, dte int) expire {
	e := expire{
		ExpireDate:   fakeDate(y, m, d),
		DaysToExpire: dte,
	}

	return e
}

func TestNewExpire(t *testing.T) {
	today := fakeDate(2015, 10, 5)

	r := refill{
		Date:     fakeDate(2015, 10, 3),
		Quantity: 80,
	}

	ratio := 3.0

	got := newExpire(today, r, ratio)

	want := fakeExpire(2015, 10, 29, 24)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("newExpire(,,) = %v; want %v",
			got,
			want,
		)
	}
}
