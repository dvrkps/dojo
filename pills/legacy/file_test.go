package main

import (
	"bytes"
	"io"
	"testing"
	"time"
)

func fakeFileContent() io.Reader {
	c := []byte(`2015-02-14,33,Aaa,1,0
//2015-02-28,62,Cardiopirin 100 mg,1
2015-02-28,27,Bbb,1
2015-02-01,109,Ccc,2,2,3`)

	return bytes.NewReader(c)
}

func TestParseFile(t *testing.T) {
	d, err := parseFile(fakeFileContent(), time.Now())
	size := len(*d)

	const wantSize = 3
	if size != wantSize || err != nil {
		t.Errorf("len(parseFile(valid, date)) = %v, %v; want %v, nil", size, err, wantSize)
	}

	s := bytes.NewReader([]byte(""))
	d, err = parseFile(s, time.Now())
	size = len(*d)

	if size != 0 || err != nil {
		t.Errorf("len(parseFile(invalid, date)) = %v, %v; want 0, nil", size, err)
	}
}
