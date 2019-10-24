package main

import "testing"

const testDataPath = "testdata/testdata.txt"

func TestLoadData(t *testing.T) {
	d, err := loadData(new(Logger), testDataPath)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if len(d) == 0 {
		t.Error("empty data")
	}
}

func TestLoadDataInvalid(t *testing.T) {
	d, err := loadData(new(Logger), "invalid")
	if err == nil {
		t.Error("nil error for invalid path")
	}
	if len(d) != 0 {
		t.Error("data must be empty")
	}
}
