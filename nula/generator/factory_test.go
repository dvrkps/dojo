package main

import (
	"strconv"
	"testing"
	"time"
)

func fakeDate(y, m, d int) time.Time {
	if y == 0 && m == 0 && d == 0 {
		return time.Time{}
	}
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func TestFactory_NewJobID(t *testing.T) {
	startEpoch := fakeDate(1970, 1, 1)
	tests := []struct {
		c    int
		now  time.Time
		sid  int
		want string
	}{
		{
			c:    1,
			now:  startEpoch,
			sid:  42,
			want: "0242",
		},
		{
			c:    factoryLimit - 1,
			now:  startEpoch,
			sid:  42,
			want: "0" + strconv.Itoa(factoryLimit) + "42",
		},
		{
			c:    factoryLimit,
			now:  startEpoch,
			sid:  42,
			want: "0142",
		},
		{
			c:    8,
			now:  fakeDate(2012, 3, 14),
			sid:  45,
			want: "1331683200000945",
		},
	}

	for _, tt := range tests {
		f := NewFactory()
		f.count = tt.c
		before := f.count
		if got := f.NewJobID(tt.now, tt.sid); got != tt.want {
			t.Errorf("(%v).NewJobID(%v) = %v; want %v",
				before,
				tt.now.Format("2006-01-02"),
				got,
				tt.want)
		}
	}

}
