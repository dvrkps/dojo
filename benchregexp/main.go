package main

import "regexp"

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
