package mecab

import "strings"

type MTable map[string]string

var (
	HebonToKunreiMap MTable
	KunreiToHebonMap MTable
	table            = MTable{
		"shi": "si",
		"chi": "ti", "tsu": "tu",
		"fu":  "hu",
		"ji":  "zi",
		"sha": "sya", "shu": "syu", "sho": "syo",
		"cha": "tya", "chu": "tyu", "cho": "tyo",
		"ja": "zya", "ju": "zyu", "jo": "zyo",
	}
)

func HebonToKunrei(word string) string {
	for k, v := range HebonToKunreiMap {
		word = strings.Replace(word, k, v, -1)
	}
	return word
}

func KunreiToHebon(word string) string {
	for k, v := range KunreiToHebonMap {
		word = strings.Replace(word, k, v, -1)
	}
	return word
}

func init() {
	HebonToKunreiMap = make(MTable)
	KunreiToHebonMap = make(MTable)

	for k, v := range table {
		HebonToKunreiMap[k] = v
		KunreiToHebonMap[v] = k
	}
}
