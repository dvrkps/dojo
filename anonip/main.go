package main

import (
	"net"
)

func AnonymizeIP(rawIP string) net.IP {
	ip := net.ParseIP(rawIP)
	if ip == nil {
		// invalid ip
		return ip
	}
	if ip.To4() != nil {
		// IPv4
		mask := net.CIDRMask(16, 32)
		return ip.Mask(mask)
	}

	mask := net.CIDRMask(48, 128)
	return ip.Mask(mask)
}
