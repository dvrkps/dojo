package main

import "fmt"

func main() {}

const (
	defaultName  = "empty"
	defaultLimit = 10
	minLimit     = 1
)

func payload(name string, max int) chan string {
	if name == "" {
		name = defaultName
	}
	if max < minLimit {
		max = defaultMax
	}
	return fmt.Sprintf("%s%d", name, max)
}
