package main

import (
	"bufio"
	"bytes"
	"os/user"
	"testing"
	"time"
)

func fakeFileContent() *bufio.Scanner {
	c := []byte(`2015-02-14,33,Aldactone 50,1,0,1,0,1,0,1
2015-02-28,62,Cardiopirin 100 mg,1
2015-02-28,73,Carvelol 12.5 mg,2
2015-02-28,27,Dualtis 1000 mg,1
2015-02-14,33,Fursemid 40 mg,0,1,0,0,0,1,0
2015-02-28,89,Gluformin 850 mg,3
2015-02-01,22,Kalinor,q,0,q,0,q,0,q
2015-02-28,36,Lotar 50 mg,1
2015-02-28,92,Preductal MR 35 mg,2
2015-02-28,61,Statex 40 mg,1
2015-02-01,109,Tyraq 25,2,2,3`)
	return bufio.NewScanner(bytes.NewReader(c))
}

func TestFilePath(t *testing.T) {
	u, _ := user.Current()
	hd := u.HomeDir
	want := hd + "/pills.txt"
	got := filePath()
	if got != want {
		t.Errorf("filePath()= %v; want %v", got, want)
	}
}

func TestFileScanner(t *testing.T) {
	if _, err := fileScanner("pills.txt"); err != nil {
		t.Errorf("fileScanner(\"pills.txt\") = _, %v; want <nil>", err)
	}
	if _, err := fileScanner("invalidpath"); err == nil {
		t.Error("fileScanner(\"invalidpath\") = _, <nil>; want error")
	}
}

func TestMidnight(t *testing.T) {
	in := time.Date(2015, 8, 11, 1, 2, 3, 4, time.UTC)
	want := time.Date(2015, 8, 11, 0, 0, 0, 0, time.UTC)
	if got := midnight(in); got != want {
		t.Errorf("midnight(%v) = %v, want %v", in, got, want)
	}
}

func TestParseFile(t *testing.T) {
	date := midnight(time.Date(2015, 8, 11, 1, 2, 3, 4, time.UTC))
	if d := parseFile(fakeFileContent(), date); len(d) != 11 {
		t.Errorf("len(parseFile(valid, date)) = %v, want 11", len(d))
	}
	s := bufio.NewScanner(bytes.NewReader([]byte("")))
	if d := parseFile(s, date); len(d) != 0 {
		t.Errorf("len(parseFile(invalid, date)) = %v, want 0", len(d))
	}
}

func TestParseFileCommentedLine(t *testing.T) {
	date := midnight(time.Date(2015, 8, 11, 1, 2, 3, 4, time.UTC))
	fileContent := []byte(`// 2015-02-28,62,Cardiopirin 100 mg,1
2015-02-28,73,Carvelol 12.5 mg,2
//2015-02-28,73,Carvelol 12.5 mg,2
 //2015-02-28,27,Dualtis 1000 mg,1`)
	s := bufio.NewScanner(bytes.NewReader(fileContent))
	if d := parseFile(s, date); len(d) != 1 {
		t.Errorf("len(parseFile(commentedlines, date)) = %v, want 1", len(d))
	}
}