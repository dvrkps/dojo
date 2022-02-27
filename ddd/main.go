package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dvrkps/dojo/ddd/real"
	"github.com/dvrkps/dojo/ddd/user"
)

func main() {
	ctx := context.Background()
	s := real.NewService()
	us, err := run(ctx, &s)
	if err != nil {
		log.Printf("run: %v", err)
		return
	}

	for _, u := range us {
		fmt.Println(u.ID, u.Name)
	}

}

func run(ctx context.Context, s user.Service) ([]user.User, error) {
	return s.Users(ctx)
}
