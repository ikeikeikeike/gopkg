package template

import "html/template"

var HelperFuncs []template.FuncMap

func init() {
	HelperFuncs = []template.FuncMap{
		{
			"toAge":              ToAge,
			"safeHTML":           SafeHTML,
			"escapeHTML":         EscapeHTML,
			"sanitizeHTML":       SanitizeHTML,
			"markdownHTML":       MarkdownHTML,
			"nl2br":              Nl2br,
			"toList":             List,
			"timeSince":          TimeSince,
			"diffLineTypeToStr":  DiffLineTypeToStr,
			"diffTypeToStr":      DiffTypeToStr,
			"sha1":               EncodeSha1,
			"toUnix":             ToUnix,
			"truncate":           Truncate,
			"toDay":              ToDay,
			"toMonth":            ToMonth,
			"toYear":             ToYear,
			"datenow":            Datenow,
			"shuffle":            Shuffle,
			"reverse":            Reverse,
			"slice":              Slice,
			"plus":               Plus,
			"autoLink":           AutoLink,
			"extractIMGs":        ExtractIMGs,
			"splitFoldl":         SplitFoldl,
			"splitFoldr":         SplitFoldr,
		},
	}
}
