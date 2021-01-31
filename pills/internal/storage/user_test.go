package storage_test

import (
	"testing"

	"github.com/dvrkps/dojo/pills/internal/storage"
)

func TestUserPath(t *testing.T) {
	t.Parallel()

	got := storage.UserPath("root", "user")

	const want = "root/pills/user"

	if got != want {
		t.Errorf("got %s; want %s", got, want)
	}
}
