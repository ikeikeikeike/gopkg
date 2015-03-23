package rdm

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestRandomNumber(t *testing.T) {
	names := []string{
		"a", "b", "c", "d",
	}

	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
	pp.Println(names[RandomNumber(0, len(names))])
}
