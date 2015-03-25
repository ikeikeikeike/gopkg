package mecab

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
import "C"

import (
	"strings"

	"github.com/gojp/kana"
)

type mecabError struct {
	message string
}

func (err mecabError) Error() (s string) {
	return err.message
}

func Parse(sentence, options string) (r []ParseResult, err error) {
	model := C.mecab_model_new2(C.CString(options))
	if model == nil {
		err = mecabError{"mecab model is not created."}
		return
	}
	mecab := C.mecab_model_new_tagger(model)
	if mecab == nil {
		err = mecabError{"mecab tagger is not created."}
		return
	}
	lattice := C.mecab_model_new_lattice(model)
	if lattice == nil {
		err = mecabError{"mecab lattice is not created."}
		return
	}

	C.mecab_lattice_set_sentence(lattice, C.CString(sentence))
	C.mecab_parse_lattice(mecab, lattice)

	lines := strings.Split(C.GoString(C.mecab_lattice_tostr(lattice)), "\n")
	for _, l := range lines {
		if strings.Index(l, "EOS") != 0 {
			if len(l) > 1 {
				r = append(r, split(l))
			}
		}
	}

	C.mecab_destroy(mecab)
	C.mecab_lattice_destroy(lattice)
	C.mecab_model_destroy(model)

	return
}

func split(line string) (r ParseResult) {
	l := strings.Split(line, "\t")
	r.Surface = l[0]
	r.Feature = l[1]

	feature := strings.Split(r.Feature, ",")
	r.Pos = feature[0]
	r.Pos1 = feature[1]
	r.Pos2 = feature[2]
	r.Pos3 = feature[3]
	r.Cform = feature[4]
	r.Ctype = feature[5]
	r.Base = feature[6]
	if len(feature) > 7 {
		r.Read = feature[7]
		r.Pron = feature[8]
		r.Romaji = kana.KanaToRomaji(r.Read)
		r.Kunrei = HebonToKunrei(r.Romaji)
		r.Hiragana = kana.RomajiToHiragana(r.Romaji)
	}
	return
}
