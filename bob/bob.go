// Package bob returns a teen, Bob's, snarky responses
package bob

import (
	"strings"
	"unicode"
)

// Hey returns teen's snarky responses
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)
	isQuestion := strings.LastIndex(trimmedRemark, "?") == (len(trimmedRemark) - 1)
	isCaps := strings.ToUpper(remark) == remark
	isWhiteSpace := len(trimmedRemark) == 0

	// to determine if a letter exists
	f := func(c rune) bool {
		return unicode.IsLetter(c)
	}

	hasLetter := strings.IndexFunc(remark, f) != -1

	if isWhiteSpace {
		return "Fine. Be that way!"
	}

	if isCaps && isQuestion && hasLetter {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion {
		return "Sure."
	}

	if isCaps && hasLetter {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
