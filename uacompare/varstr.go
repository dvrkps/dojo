package main

import (
	varstrPkg "github.com/varstr/uaparser"
)

const varstrName = "varstr"

func varstrParse(pkg *uaPackage, in string) *uaResult {
	ua := varstrPkg.Parse(in)
	if ua == nil {
		return &uaResult{}
	}

	osName, osVersion := varstrNameVersion(ua.OS)
	browserName, browserVersion := varstrNameVersion(ua.Browser)

	mobile, tablet := varstrDeviceType(ua)

	r := uaResult{
		os:             osName,
		osVersion:      osVersion,
		browser:        browserName,
		browserVersion: browserVersion,
		mobile:         mobile,
		tablet:         tablet,
	}
	if r.Valid() {
		pkg.Inc()
	}
	return &r
}

func varstrNameVersion(ii *varstrPkg.InfoItem) (string, string) {
	var n, v string
	if ii != nil {
		n = ii.Name
		v = ii.Version
	}
	return n, v
}

func varstrDeviceType(ua *varstrPkg.UAInfo) (bool, bool) {
	var mobile, tablet bool
	if ua.DeviceType != nil {
		switch ua.DeviceType.Name {
		case "Phone":
			mobile = true
		case "Tablet":
			tablet = true
		}
	}
	return mobile, tablet
}
