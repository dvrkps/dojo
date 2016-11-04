package availability

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {

	tests := []struct {
		typ  int
		want *Availability
		ok   bool
	}{

		{
			typ:  TimeType,
			want: &Availability{typ: TimeType},
			ok:   true,
		},

		{
			typ:  AggregateType,
			want: &Availability{typ: AggregateType},
			ok:   true,
		},

		{
			typ:  TimeType - 1,
			want: nil,
			ok:   false,
		},

		{
			typ:  AggregateType + 1,
			want: nil,
			ok:   false,
		},
	}

	for _, tt := range tests {

		got, err := New(tt.typ)

		if !tt.ok {

			if got != nil || err == nil {

				t.Errorf("New(%d) = %+v, nil; want <nil>, <error>",
					tt.typ, tt.want)
			}

			continue
		}

		if !reflect.DeepEqual(got, tt.want) || err != nil {

			t.Errorf("New(%d) = %v, %v; want %v, <error>",
				tt.typ, got, err, tt.want)
		}

	}
}
