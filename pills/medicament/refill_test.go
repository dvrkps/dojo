package medicament

import (
	"reflect"
	"testing"
)

func TestNewRefill(t *testing.T) {
	date := []byte("2015-09-29")
	quantity := []byte("12.6")
	want := refill{
		Date:     fakeDate(2015, 9, 29),
		Quantity: 12.6,
	}
	got, err := newRefill(date, quantity)
	if !reflect.DeepEqual(got, want) || err != nil {
		t.Errorf("newRefill(%q,%q) = %v, %v; want %v, nil",
			date,
			quantity,
			got,
			err,
			want,
		)
	}
}

func TestNewRefill_errors(t *testing.T) {
	tests := []struct {
		date, quantity []byte
	}{
		{[]byte("abc"), []byte("123")},
		{[]byte("2015-09-29"), []byte("abc")},
		{[]byte("2015-09-29"), []byte("-14.876")},
	}
	for _, tt := range tests {
		_, err := newRefill(tt.date, tt.quantity)
		if err == nil {
			t.Errorf("newRefill(%q,%q) = _, %v; want empty, error",
				tt.date,
				tt.quantity,
				err,
			)
		}
	}
}
