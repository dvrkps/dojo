package medicament

import (
	"reflect"
	"testing"
)

func TestDosageCalcRatio(t *testing.T) {
	tests := []struct {
		in   []float64
		want float64
	}{
		{[]float64{0, 1}, 0.5},
		{[]float64{0, 0.25, 0, 0.25}, 0.125},
		{[]float64{2, 2, 3}, 2.3333333333333335},
		{[]float64{1, 0, 1, 0, 1, 0, 1}, 0.5714285714285714},
		{[]float64{}, 0},
		{[]float64{-3, -2, -1}, 0},
	}
	for _, tt := range tests {
		if got := calcDosageRatio(tt.in...); got != tt.want {
			t.Errorf("calcDosageRatio(%v) = %v; want %v",
				tt.in,
				got,
				tt.want,
			)
		}
	}
}

func TestDosage_Ratio(t *testing.T) {
	want := 5.0
	in := Dosage{ratio: want}

	if got := in.Ratio(); got != want {
		t.Errorf("Ratio() = %v; want %v", got, want)
	}
}

func TestDosage_String(t *testing.T) {
	in := Dosage{doses: []float64{1, 0, 0.25, 0.5, 2.0, 4}}
	want := "1 - q h 2 4"

	if got := in.String(); got != want {
		t.Errorf("String() = %q; want %q", got, want)
	}
}

func TestNewDosage(t *testing.T) {
	in := [][]byte{
		[]byte("1"),
		[]byte("h"),
		[]byte("2"),
		[]byte("q"),
		[]byte("3"),
	}
	want := Dosage{doses: []float64{1, 0.5, 2, 0.25, 3}, ratio: 1.35}

	if got, err := newDosage(in...); !reflect.DeepEqual(got, want) || err != nil {
		t.Errorf("newDosage(%q) = %v, %v; want %v, <nil>",
			in,
			got,
			err,
			want,
		)
	}
}

func TestNewDosage_empty(t *testing.T) {
	in := [][]byte{}
	if _, err := newDosage(in...); err == nil {
		t.Errorf("newDosage(%q) = _, %v; want empty, error",
			in,
			err,
		)
	}
}

func TestNewDosage_error(t *testing.T) {
	in := [][]byte{
		[]byte("5"),
		[]byte("m"),
		[]byte(" 2a"),
	}
	if _, err := newDosage(in...); err == nil {
		t.Errorf("newDosage(%q) = _, %v; want empty, error",
			in,
			err,
		)
	}
}

func TestNewDose(t *testing.T) {
	tests := []struct {
		in   []byte
		want float64
	}{
		{[]byte("h"), 0.5},
		{[]byte("q"), 0.25},
		{[]byte("1"), 1},
		{[]byte("0.50"), 0.5},
		{[]byte("2.35"), 2.35},
	}
	for _, tt := range tests {
		if got, err := newDose(tt.in); got != tt.want || err != nil {
			t.Errorf("newDose(%q) = %v, %v; want %v, <nil>",
				tt.in,
				got,
				err,
				tt.want,
			)
		}
	}
}

func TestNewDose_errors(t *testing.T) {
	tests := [][]byte{
		[]byte(".2"),
		[]byte("3."),
		[]byte(" 0.10"),
		[]byte("a "),
		[]byte("b"),
		[]byte("-42"),
		[]byte(""),
	}
	for _, tt := range tests {
		if got, err := newDose(tt); got != 0 || err == nil {
			t.Errorf("newDose(%q) = %v,<nil>; want 0, error", tt, got)
		}
	}
}
