package main

import (
	"bytes"
	"io"
	"testing"
	"time"
)

func fakeFileContent() io.Reader {
	c := []byte(`2015-02-14,33,Aldactone 50,1,0,1,0,1,0,1
2015-02-28,62,Cardiopirin 100 mg,1
2015-02-28,73,Carvelol 12.5 mg,2
2015-02-28,27,Dualtis 1000 mg,1
2015-02-14,33,Fursemid 40 mg,0,1,0,0,0,1,0
2015-02-28,89,Gluformin 850 mg,3
2015-02-01,22,Kalinor,0.25,0,0.25,0,0.25,0,0.25
2015-02-28,36,Lotar 50 mg,1
2015-02-28,92,Preductal MR 35 mg,2
2015-02-28,61,Statex 40 mg,1
2015-02-01,109,Tyraq 25,2,2,3`)

	return bytes.NewReader(c)
}

func testDate() time.Time {
	return time.Date(2015, 8, 11, 0, 0, 0, 0, time.UTC)
}

func TestParseFile(t *testing.T) {
	d, err := parseFile(fakeFileContent(), testDate())
	size := len(*d)

	if size != 11 || err != nil {
		t.Errorf("len(parseFile(valid, date)) = %v, %v; want 11, nil", size, err)
	}

	s := bytes.NewReader([]byte(""))
	d, err = parseFile(s, testDate())
	size = len(*d)

	if size != 0 || err != nil {
		t.Errorf("len(parseFile(invalid, date)) = %v, %v; want 0, nil", size, err)
	}
}

func TestParseFileCommentedLine(t *testing.T) {
	fileContent := []byte(`// 2015-02-28,62,Cardiopirin 100 mg,1
2015-02-28,73,Carvelol 12.5 mg,2
//2015-02-28,73,Carvelol 12.5 mg,2
 //2015-02-28,27,Dualtis 1000 mg,1`)
	s := bytes.NewReader(fileContent)
	d, err := parseFile(s, testDate())
	size := len(*d)

	if size != 1 || err != nil {
		t.Errorf("len(parseFile(commentedlines, date)) = %v, %v; want 0, nil", size, err)
	}
}
