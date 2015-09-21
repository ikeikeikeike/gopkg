package template

import "strings"

func SplitFoldl(s, sep string) string {
	return strings.Split(s, sep)[0]
}

func SplitFoldr(s, sep string) string {
	return strings.Split(s, sep)[1]
}
