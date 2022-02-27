package real

import (
	"context"

	"github.com/dvrkps/dojo/ddd/user"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s *Service) Users(ctx context.Context) ([]user.User, error) {
	r := []user.User{
		{ID: 1, Name: "Real one"},
		{ID: 2, Name: "Real two"},
	}
	return r, nil
}
