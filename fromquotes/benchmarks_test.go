package main

func benchCases() map[string][]string {
	m := map[string][]string{
		"short": short(),
		"long":  long(),
	}
	return m
}
