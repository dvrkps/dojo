package availability

import (
	"reflect"
	"testing"
)

func TestNewTimeBased(t *testing.T) {

	tests := []struct {
		want *TimeBased
	}{

		{
			want: &TimeBased{yearDowntime: yearNs},
		},
	}

	for _, tt := range tests {

		if got := NewTimeBased(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewTimeBased() = %v; want %v",
				got, tt.want)
		}

	}
}
