package str

import "strings"

func Clean(s string) string {
	return strings.Trim(strings.Trim(s, " "), "　")
}
