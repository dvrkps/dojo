package mock

import (
	"context"

	"github.com/dvrkps/dojo/ddd/user"
)

type Service struct {
	UsersFn func(context.Context) ([]user.User, error)
}

func (s *Service) Users(ctx context.Context) ([]user.User, error) {
	return s.UsersFn(ctx)
}

func Users() []user.User {
	r := []user.User{
		{ID: 41, Name: "Mock one"},
		{ID: 42, Name: "Mock two"},
	}

	return r
}
