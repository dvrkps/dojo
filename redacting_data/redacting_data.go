package main

import "fmt"

// Config holds configuration data.
type Config struct {
	Value  int
	Secret string
}

func (c *Config) String() string {
	type clone Config
	cc := clone(*c) // avoid recursion
	cc.Secret = "(REDACTED)"
	return fmt.Sprintf("%+v", cc)
}

func main() {
	c := &Config{
		Value:  42,
		Secret: "Enigma",
	}

	fmt.Println(c, c.Secret)

}
