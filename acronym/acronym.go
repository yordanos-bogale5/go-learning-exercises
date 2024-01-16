// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym converts long phrases into their acronyms
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes string s and converts it into an acronym with to
// ignore certain chracters
func Abbreviate(s string) string {
	re := regexp.MustCompile(`[A-Za-z']+`)
	matches := re.FindAll([]byte(s), -1)
	var as []byte

	for _, word := range matches {
		as = append(as, word[0])
	}

	return strings.ToUpper(string(as))
}
