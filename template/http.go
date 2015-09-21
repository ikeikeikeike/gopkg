package template

import "strings"

func ToCanonical(name string) string {
	name = strings.Replace(name, "!", "！", -1)
	name = strings.Replace(name, "`", "｀", -1)
	name = strings.Replace(name, "~", "〜", -1)
	name = strings.Replace(name, "'", "’", -1)
	name = strings.Replace(name, `"`, "”", -1)
	name = strings.Replace(name, `[`, "「", -1)
	name = strings.Replace(name, `]`, "」", -1)
	name = strings.Replace(name, `(`, "（", -1)
	name = strings.Replace(name, `)`, "）", -1)
	name = strings.Replace(name, `{`, "｛", -1)
	name = strings.Replace(name, `}`, "｝", -1)
	name = strings.Replace(name, `;`, "；", -1)
	name = strings.Replace(name, `:`, "：", -1)
	name = strings.Replace(name, `<`, "＜", -1)
	name = strings.Replace(name, `>`, "＞", -1)
	name = strings.Replace(name, `,`, "、", -1)
	name = strings.Replace(name, `.`, "．", -1)
	name = strings.Replace(name, `*`, "＊", -1)
	name = strings.Replace(name, `&`, "＆", -1)
	name = strings.Replace(name, `^`, "＾", -1)
	name = strings.Replace(name, `$`, "＄", -1)
	name = strings.Replace(name, `@`, "＠", -1)
	name = strings.Replace(name, "\u3000", " ", -1)
	name = strings.Replace(name, "　", " ", -1)
	name = strings.Replace(name, " ", "_", -1)
	return name
}
