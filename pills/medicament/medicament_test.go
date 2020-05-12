package medicament

import (
	"reflect"
	"testing"
	"time"
)

func fakeDate(y, m, d int) time.Time {
	if y == 0 && m == 0 && d == 0 {
		var t time.Time

		return t
	}

	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func TestMedicament_String(t *testing.T) {
	m := Medicament{
		Name:   "abc",
		dosage: dosage{0, 1},
		expire: expire{
			ExpireDate:   fakeDate(2015, 3, 1),
			DaysToExpire: 0,
		},
	}

	want := "  0     1.3. Sun   abc [0 1]"
	if got := m.String(); got != want {
		t.Errorf("String() = %v; want %v", got, want)
	}
}

func TestCompare(t *testing.T) {
	a := Medicament{
		expire: expire{
			DaysToExpire: 10,
		},
	}

	b := Medicament{
		expire: expire{
			DaysToExpire: 12,
		},
	}

	if got, want := Compare(a, b), true; got != want {
		t.Errorf("Compare(m1 < m2) = %v; want %v", got, want)
	}

	if got, want := Compare(b, a), false; got != want {
		t.Errorf("Compare(m1 > m2) = %v; want %v", got, want)
	}

	if got, want := Compare(a, a), false; got != want {
		t.Errorf("Compare(m1 == m2) = %v; want %v", got, want)
	}
}

func TestMidnight(t *testing.T) {
	in := time.Date(2015, 9, 29, 1, 2, 3, 4, time.UTC)
	want := time.Date(2015, 9, 29, 0, 0, 0, 0, time.UTC)

	if got := midnight(in); got != want {
		t.Errorf("midnight(%v) = %v, want %v", in, got, want)
	}
}

func TestNew(t *testing.T) {
	today := fakeDate(2015, 3, 1)
	in := []byte("2015-02-24,12,abc,0,1")
	want := Medicament{
		refill: refill{
			Date:     fakeDate(2015, 2, 24),
			Quantity: 12,
		},
		Name:   "abc",
		dosage: dosage{0, 1},
		expire: expire{
			ExpireDate:   fakeDate(2015, 3, 20),
			DaysToExpire: 19,
		},
	}

	if got, err := New(today, in); !reflect.DeepEqual(got, want) || err != nil {
		t.Errorf("New(,) = %+v, %v; want %+v, nil", got, err, want)
	}
}

func TestNew_errors(t *testing.T) {
	today := fakeDate(2015, 3, 1)
	tests := []struct {
		in []byte
	}{
		{[]byte("2015-02-24,12,def,-99,xyz")},
		{[]byte("2015-02-24,12,,0,1")},
		{[]byte("wrongdate,12,,0,1")},
		{[]byte("")},
	}

	for _, tt := range tests {
		if got, err := New(today, tt.in); !reflect.DeepEqual(got, Medicament{}) ||
			err == nil {
			t.Errorf("New(today, %q) = %+v, %v; want empty, error", tt.in, got, err)
		}
	}
}

func TestParseName(t *testing.T) {
	in := []byte("Aldactone")
	want := "Aldactone"

	if got, err := parseName(in); got != want || err != nil {
		t.Errorf("parseName(%q) = %v, %v; want %v, nil",
			in,
			got,
			err,
			want,
		)
	}
}

func TestParseName_empty(t *testing.T) {
	in := []byte{}

	if _, err := parseName(in); err == nil {
		t.Errorf("parseName(%q) = _, %v; want _, error",
			in,
			err,
		)
	}
}
