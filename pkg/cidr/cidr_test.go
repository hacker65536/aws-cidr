package cidr

import (
	"testing"
)

var (
	PrdCidr = "10.64.0.0/11"
	StgCidr = "10.128.0.0/11"
	DevCidr = "10.192.0.0/11"
)

func TestIsIPInCIDR(t *testing.T) {
	ip := "10.64.0.0"
	cidr := "10.64.0.0/11"
	ipContained, err := IsIPInCIDR(ip, cidr)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if ipContained {
		t.Logf("IP %v contained in CIDR %v", ip, cidr)
	}
}

func TestIsCIDRInCIDR(t *testing.T) {
	smallCIDR := "10.156.16.0/20"
	largeCIDR := StgCidr
	cidrContained, err := IsCIDRInCIDR(smallCIDR, largeCIDR)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if cidrContained {
		t.Logf("CIDR %v contained in CIDR %v", smallCIDR, largeCIDR)
	}
}
