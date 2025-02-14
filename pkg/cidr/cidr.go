package cidr

import (
	"fmt"
	"net"
)

// isIPInCIDR checks if the given IP address is within the specified CIDR block.
func IsIPInCIDR(ipStr, cidrStr string) (bool, error) {
	// Parse the IP address.
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false, fmt.Errorf("invalid IP address: %s", ipStr)
	}

	// Parse the CIDR block.
	_, network, err := net.ParseCIDR(cidrStr)
	if err != nil {
		return false, fmt.Errorf("invalid CIDR block: %s", cidrStr)
	}

	// Check if the IP address is within the CIDR block.
	return network.Contains(ip), nil
}

// isCIDRInCIDR checks if the given CIDR block is within the specified CIDR block.
func IsCIDRInCIDR(smallCIDR, largeCIDR string) (bool, error) {
	// Parse the small CIDR block.
	_, smallNetwork, err := net.ParseCIDR(smallCIDR)
	if err != nil {
		return false, fmt.Errorf("invalid small CIDR block: %s", smallCIDR)
	}

	// Parse the large CIDR block.
	_, largeNetwork, err := net.ParseCIDR(largeCIDR)
	if err != nil {
		return false, fmt.Errorf("invalid large CIDR block: %s", largeCIDR)
	}

	// Check if the small CIDR block is within the large CIDR block.
	return largeNetwork.Contains(smallNetwork.IP), nil
}

/*
func main() {
	// Example usage for IP address
	ip := "192.168.1.5"
	cidr := "192.168.1.0/24"

	ipContained, err := isIPInCIDR(ip, cidr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if ipContained {
		fmt.Printf("The IP address %s is within the CIDR block %s.\n", ip, cidr)
	} else {
		fmt.Printf("The IP address %s is not within the CIDR block %s.\n", ip, cidr)
	}

	// Example usage for CIDR block
	smallCIDR := "192.168.1.128/25"
	largeCIDR := "192.168.1.0/24"

	cidrContained, err := isCIDRInCIDR(smallCIDR, largeCIDR)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if cidrContained {
		fmt.Printf("The CIDR block %s is within the CIDR block %s.\n", smallCIDR, largeCIDR)
	} else {
		fmt.Printf("The CIDR block %s is not within the CIDR block %s.\n", smallCIDR, largeCIDR)
	}
}
*/
