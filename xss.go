package main

import (
	"regexp"
)

var replaceWhitespaces = regexp.MustCompile(`\s\s+`)
var blacklistStripping = regexp.MustCompile(`[\p{Me}\p{C}<>=;(){}\[\]?&]`)

func xss(s string) string {
	s = blacklistStripping.ReplaceAllString(s, " ")
	s = replaceWhitespaces.ReplaceAllString(s, " ")
	return s
}
