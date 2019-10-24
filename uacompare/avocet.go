package main

import (
	"fmt"
	"strings"

	avocetPkg "github.com/avct/uasurfer"
)

const avocetName = "avocet"

func avocetParse(pkg *uaPackage, in string) *uaResult {
	ua := avocetPkg.Parse(in)
	if ua == nil {
		return &uaResult{}
	}

	osName, osVersion := avocetNameVersion(ua.OS.Name.String(), ua.OS.Version)
	browserName, browserVersion := avocetNameVersion(ua.Browser.Name.String(), ua.Browser.Version)
	mobile, tablet := avocetDeviceType(ua)

	osName = fixUnknown(osName)
	osVersion = fixUnknown(osVersion)
	browserName = fixUnknown(browserName)
	browserVersion = fixUnknown(browserVersion)

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

func avocetNameVersion(rawName string, rawVersion avocetPkg.Version) (string, string) {
	const (
		osPrefix      = "OS"
		browserPrefix = "Browser"
	)

	var name string
	if strings.HasPrefix(rawName, osPrefix) {
		name = strings.TrimPrefix(rawName, osPrefix)
	}
	if strings.HasPrefix(rawName, browserPrefix) {
		name = strings.TrimPrefix(rawName, browserPrefix)
	}

	version := fmt.Sprintf("%d.%d.%d", rawVersion.Major, rawVersion.Minor, rawVersion.Patch)

	return name, version
}

func avocetDeviceType(ua *avocetPkg.UserAgent) (bool, bool) {
	var mobile, tablet bool
	if ua.DeviceType == avocetPkg.DevicePhone {
		mobile = true
	}
	if ua.DeviceType == avocetPkg.DeviceTablet {
		tablet = true
	}
	return mobile, tablet
}
