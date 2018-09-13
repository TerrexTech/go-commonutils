package commonutil

import (
	"strings"
)

// ParseHosts extracts hosts from specified string.
// For example, a string having the value:
//  "192.168.1.1:8080, 172.17.1.1,10.15.1.1"
// will be extracted into a slice containing data as:
//  ["192.168.1.1:8080", "172.17.1.1", "10.15.1.1"]
// Any spaces after any comma are automatically removed.
// Returns blank slice if the string is blank.
func ParseHosts(s string) *[]string {
	if s != "" {
		// Remove spaces after commas
		s = strings.Replace(s, ", ", ",", -1)
		hosts := strings.Split(s, ",")
		return &hosts
	}
	return &[]string{}
}

// StandardizeSpaces removes all extra whitespaces and newlines from string
func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
