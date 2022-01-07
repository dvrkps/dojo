package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

func main() {
	err := run()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func run() error {
	key, err := loadSecretKey()
	if err != nil {
		return fmt.Errorf("load secret key: %v", err)

	}

	ctx := context.Background()

	gc := githubClient(ctx, key)

	// list all repositories for the authenticated user
	repos, _, err := gc.Repositories.List(ctx, "", nil)
	if err != nil {
		return err
	}

	for i, r := range repos {
		fmt.Printf("%v %v\n", i, r.GetName())
	}

	return nil
}

func loadSecretKey() (string, error) {
	b, err := os.ReadFile("secret.txt")
	return string(bytes.TrimSpace(b)), err
}

func githubClient(ctx context.Context, key string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: key},
	)

	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}
