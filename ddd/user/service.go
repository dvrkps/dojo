package user

import "context"

type Service interface {
	Users(ctx context.Context) ([]User, error)
}
