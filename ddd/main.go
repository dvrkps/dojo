package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dvrkps/dojo/ddd/real"
)

func main() {
	s := real.NewService()

	us, err := s.Users(context.Background())
	if err != nil {
		log.Printf("users: %v", err)
		return
	}

	for _, u := range us {
		fmt.Println(u.ID, u.Name)
	}
}
