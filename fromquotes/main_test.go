package main

type testCase struct {
	in  []string
	out []string
}

func testCases() map[string]testCase {
	m := map[string]testCase{
		"short": []string{".hi", "\"My", "name", "is", "Omar\"", "123", "\"123\""},
		"long":  []string{"\".hi", "I'm", "the", "real", "Slim", "Shady", "\"My", "name", "is", "Omar\"", "hello", "world", "\"123\"", "a"},
	}
	return m
}

var result []string
