package main

import (
	"context"
	"testing"

	"github.com/dvrkps/dojo/ddd/real/mock"
	"github.com/dvrkps/dojo/ddd/user"
)

func TestRun(t *testing.T) {
	s := mock.Service{
		UsersFn: func(ctx context.Context) ([]user.User, error) {
			return mock.Users(), nil
		},
	}

	got, err := run(context.Background(), &s)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	want := mock.Users()
	if !compare(got, want) {
		t.Error("got not equal want")
	}
}

func compare(got, want []user.User) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
