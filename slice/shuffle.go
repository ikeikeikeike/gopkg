package slice

import (
	"math/rand"
	"time"
)

func Shuffle(ref interface{}) {
	switch ref.(type) {
	case []int:
		ary := ref.([]int)
		for i, v := range ShuffleInt(ary) {
			ary[i] = v
		}
	case []string:
		ary := ref.([]string)
		for i, v := range ShuffleString(ary) {
			ary[i] = v
		}
	}

}

func ShuffleInt(src []int) []int {
	dest := make([]int, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func ShuffleString(src []string) []string {
	dest := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}
