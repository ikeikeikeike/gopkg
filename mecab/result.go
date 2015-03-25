package mecab

type ParseResult struct {
	Surface  string
	Feature  string
	Pos      string
	Pos1     string
	Pos2     string
	Pos3     string
	Cform    string
	Ctype    string
	Base     string
	Read     string
	Pron     string
	Romaji   string
	Kunrei   string
	Hiragana string
}
