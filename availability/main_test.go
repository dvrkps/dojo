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

func TestByRequests_SetPercent(t *testing.T) {
	tests := []struct {
		in float64
		ok bool
	}{
		{
			in: 0,
			ok: false,
		},

		{
			in: 100,
			ok: false,
		},

		{
			in: 0.1,
			ok: true,
		},

		{
			in: 99.999,
			ok: true,
		},
	}

	for _, tt := range tests {
		br := NewByRequests()
		err := br.SetPercent(tt.in)

		if !tt.ok {
			if br.percent != 0 || err == nil {
				t.Errorf("SetPercent(%v) = %v; want <error>; br: %+v",
					tt.in, err, br)
			}
			continue
		}

		if br.percent != tt.in || err != nil {
			t.Errorf("SetPercent(%v) = %v; want <nil>; br: %+v",
				tt.in, err, br)
		}
	}
}

func TestByRequests_SetSuccess(t *testing.T) {
	tests := []struct {
		in int
		ok bool
	}{
		{
			in: 0,
			ok: false,
		},

		{
			in: 1,
			ok: true,
		},
	}

	for _, tt := range tests {
		br := NewByRequests()
		err := br.SetSuccess(tt.in)

		if !tt.ok {
			if br.success != 0 || err == nil {
				t.Errorf("SetSuccess(%v) = %v; want <error>; br: %+v",
					tt.in, err, br)
			}
			continue
		}

		if br.success != tt.in || err != nil {
			t.Errorf("SetSuccess(%v) = %v; want <nil>; br: %+v",
				tt.in, err, br)
		}
	}
}

func TestByRequests_SetTotal(t *testing.T) {
	tests := []struct {
		in int
		ok bool
	}{
		{
			in: 0,
			ok: false,
		},

		{
			in: 1,
			ok: true,
		},
	}

	for _, tt := range tests {
		br := NewByRequests()
		err := br.SetTotal(tt.in)

		if !tt.ok {
			if br.success != 0 || err == nil {
				t.Errorf("SetTotal(%v) = %v; want <error>; br: %+v",
					tt.in, err, br)
			}
			continue
		}

		if br.total != tt.in || err != nil {
			t.Errorf("SetTotal(%v) = %v; want <nil>; br: %+v",
				tt.in, err, br)
		}
	}
}
