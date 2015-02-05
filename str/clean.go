package str

import "strings"

func Clean(s string) string {
	return strings.TrimSpace(strings.Trim(strings.Trim(s, " "), "　"))
}
