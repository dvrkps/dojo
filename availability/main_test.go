package availability

import (
	"reflect"
	"testing"
)

func TestNewByTime(t *testing.T) {

	tests := []struct {
		want *ByTime
	}{

		{
			want: &ByTime{downtime: yearNs},
		},
	}

	for _, tt := range tests {

		if got := NewByTime(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewByTime() = %v; want %v",
				got, tt.want)
		}

	}
}

func TestNewByRequests(t *testing.T) {

	tests := []struct {
		want *ByRequests
	}{

		{
			want: &ByRequests{},
		},
	}

	for _, tt := range tests {

		if got := NewByRequests(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewByRequests() = %v; want %v",
				got, tt.want)
		}

	}
}
