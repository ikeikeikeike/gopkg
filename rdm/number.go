package rdm

import (
	"math/rand"
	"time"
)

func RandomNumber(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
