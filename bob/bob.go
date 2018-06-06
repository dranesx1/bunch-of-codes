// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	alphaRegex := regexp.MustCompile("[a-zA-Z]")

	question := strings.HasSuffix(remark, "?")
	containsLetter := alphaRegex.MatchString(remark)
	AllUppercase := strings.ToUpper(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}
	if remark == AllUppercase && containsLetter && !question {
		return "Whoa, chill out!"
	} else if remark == AllUppercase && containsLetter && question {
		return "Calm down, I know what I'm doing!"
	} else if question {
		return "Sure."
	}
	return "Whatever."
}
