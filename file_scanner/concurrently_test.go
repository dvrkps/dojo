package main

import "testing"

func TestScanConcurrently(t *testing.T) {
	testScan(t, scanConcurrently)
}
