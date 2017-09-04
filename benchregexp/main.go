package main

import (
	"regexp"
	"strconv"
	"strings"
)

func main() {}

func withRegexp(query map[string]string) bool {
	for key := range query {
		a, _ := regexp.MatchString("(?i)^a[0-9]{1,3}$", key)
		if a {
			return true
		}
		pv, _ := regexp.MatchString("(?i)^pv[0-9]{1,3}$", key)
		if pv {
			return true
		}
		pu, _ := regexp.MatchString("(?i)^pu[0-9]{1,3}$", key)
		if pu {
			return true
		}
		sv, _ := regexp.MatchString("(?i)^sv[0-9]{1,3}$", key)
		if sv {
			return true
		}
	}
	return false
}

func noooRegexp(query map[string]string) bool {
	for key := range query {
		key := strings.TrimSpace(key)
		a := isValid(key, "a")
		if a {
			return true
		}
		pv := isValid(key, "pv")
		if pv {
			return true
		}
		pu := isValid(key, "pu")
		if pu {
			return true
		}
		sv := isValid(key, "sv")
		if sv {
			return true
		}
	}
	return false
}

func isValid(key, prefix string) bool {
	if strings.HasPrefix(key, prefix) {
		rem := key[len(prefix):]
		if _, err := strconv.ParseInt(rem, 10, 64); err != nil {
			return false
		}
		return true
	}
	return false
}
