package medicament

import (
	"reflect"
	"testing"
)

func TestNewDosage(t *testing.T) {
	in := [][]byte{
		[]byte("1"),
		[]byte("0.5"),
		[]byte("2"),
		[]byte("0.25"),
		[]byte("3"),
	}
	want := dosage{1, 0.5, 2, 0.25, 3}

	if got, err := newDosage(in...); !reflect.DeepEqual(got, want) || err != nil {
		t.Errorf("newDosage(%q) = %v, %v; want %v, <nil>",
			in,
			got,
			err,
			want,
		)
	}
}

/*
func TestNewDose(t *testing.T) {
	tests := []struct {
		in   []byte
		want float64
	}{
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
		//[]byte(".2"),
		//[]byte("3."),
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
*/
