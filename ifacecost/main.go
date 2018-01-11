package main

import (
	"fmt"
	"strings"
)

type words struct {
	sentence string
}

func (w words) split() []string {
	return strings.Fields(w.sentence)
}

type spliter interface {
	split() []string
}

func typeSplit(w words) []string {
	return w.split()
}

func ifaceSplit(s spliter) []string {
	return s.split()
}

func main() {
	w := words{sentence: "aaa bbb ccc"}

	ts := typeSplit(w)
	fmt.Println(ts)

	is := ifaceSplit(w)
	fmt.Println(is)
}
