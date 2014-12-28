package accessctl

import "regexp"

var (
	UARegex string = "(Mozilla|Opera)"
	Re      *regexp.Regexp
)

func IsUA(ua string) bool {
	return Re.MatchString(ua)
}

func init() {
	Re, _ = regexp.Compile(UARegex)
}
