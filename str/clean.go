package str

import (
	"regexp"
	"strings"
)

func Clean(s string) string {
	return strings.TrimSpace(strings.Trim(strings.Trim(s, " "), "　"))
}

func MustClean(s string) string {
	s = strings.Replace(Clean(s), "\u3000", "", -1)
	return regexp.MustCompile(`\s{2,}`).ReplaceAllString(s, "")
}
