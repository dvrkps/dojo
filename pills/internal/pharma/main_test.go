package pharma

import (
	"io"
	"strings"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	today := time.Date(2020, 6, 8, 1, 2, 3, 4, time.UTC)
	all, err := All(fakeFileContent(), today)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	got := len(all)
	const want = 3
	if got != want {
		t.Errorf("len got %v; want %v", got, want)
	}

}

func fakeFileContent() io.Reader {
	const newline = "\n"
	const c = "2020-06-01,20,Dd,1\n" +
		"//2020-06-01,31,Cc,1\n" +
		"2020-06-01,19,Bb,1\n" +
		"2020-06-01,18,Aa,1\n"

	return strings.NewReader(c)
}
