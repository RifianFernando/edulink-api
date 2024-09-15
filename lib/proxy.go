package lib

import "strings"

// Utility function to check if ALLOW_ORIGIN is a valid IP
func IsValidIP(ip string) bool {
	// Basic check if the string represents an IP address
	return strings.Count(ip, ".") == 3 || strings.Count(ip, ":") > 0
}

func HandleTrustedProxies(allowOrigin string) []string {
	// Handle trusted proxies based on ALLOW_ORIGIN
	var trustedProxies []string
	if strings.Contains(allowOrigin, "localhost") || allowOrigin == "127.0.0.1" {
		trustedProxies = []string{"127.0.0.1/32"}
	} else if IsValidIP(allowOrigin) {
		// If allowOrigin is a valid IP
		trustedProxies = []string{allowOrigin + "/32"}
	} else {
		// Default fallback if allowOrigin is not a valid IP or localhost
		trustedProxies = []string{"0.0.0.0/0"}
	}

	return trustedProxies
}
