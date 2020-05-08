package main

import (
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

func fakePills() [][]byte {
	r := [][]byte{
		[]byte("2015-02-14,33,Aldactone 50,1,0,1,0,1,0,1"),
		[]byte("2015-02-28,62,Cardiopirin 100 mg,1"),
		[]byte("2015-02-28,73,Carvelol 12.5 mg,2"),
		[]byte("2015-02-28,27,Dualtis 1000 mg,1"),
		[]byte("2015-02-14,33,Fursemid 40 mg,0,1,0,0,0,1,0"),
		[]byte("2015-02-28,89,Gluformin 850 mg,3"),
		[]byte("2015-02-01,22,Kalinor,q,0,q,0,q,0,q"),
		[]byte("2015-02-28,36,Lotar 50 mg,1"),
		[]byte("2015-02-28,92,Preductal MR 35 mg,2"),
		[]byte("2015-02-28,61,Statex 40 mg,1"),
		[]byte("2015-02-01,109,Tyraq 25,2,2,3"),
	}

	return r
}

func TestData_Add(t *testing.T) {
	today := fakeDate(2015, 3, 1)
	f := fakePills()
	d := Data{}
	var err error

	for _, p := range f {
		d, err = d.Add(p, today)
		if err != nil {
			t.Fatalf("add err: %v", err)
		}
	}

	if got, want := len(f), len(d); got != want {
		t.Errorf("len(Data) = %v; want %v", got, want)
	}

	// invalid pill
	d = Data{}
	d, err = d.Add([]byte(""), today)
	if err == nil {
		t.Fatal("add empty pill must be error")
	}
}

func TestData_String(t *testing.T) {
	today := fakeDate(2015, 3, 1)
	f := fakePills()
	d := Data{}
	var err error

	for _, p := range f {
		d, err = d.Add(p, today)
		if err != nil {
			t.Fatalf("add err: %v", err)
		}
	}

	want := 451
	if got := len(d.String()); got != want {
		t.Errorf("len(String()) = %v, want %v", got, want)
	}
}

func TestSortData(t *testing.T) {
	today := fakeDate(2015, 3, 1)
	f := fakePills()
	d := Data{}
	var err error

	for _, p := range f {
		d, err = d.Add(p, today)
		if err != nil {
			t.Fatalf("add err: %v", err)
		}
	}

	d = sortData(d)

	d0 := d[0].DaysToExpire
	d10 := d[10].DaysToExpire

	if d0 != 18 || d10 != 126 {
		t.Error("Data sort order not valid")
	}
}
