package main

import (
	mileusnaPkg "github.com/mileusna/useragent"
)

const mileusnaName = "mileusna"

func mileusnaParse(pkg *uaPackage, in string) *uaResult {
	ua := mileusnaPkg.Parse(in)
	r := uaResult{
		os:             ua.OS,
		osVersion:      ua.OSVersion,
		browser:        ua.Name,
		browserVersion: ua.Version,
		device:         ua.Device,
		mobile:         ua.Mobile,
		tablet:         ua.Tablet,
	}
	if r.Valid() {
		pkg.Inc()
	}
	return &r
}
