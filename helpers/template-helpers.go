package helpers

import (
	"html/template"
	"time"
)

var TemplateHelpers = template.FuncMap{
	"raw":      htmlRaw,
	"mbsubstr": mbsubstr,
	"datestr":  datestr,
}

func htmlRaw(html string) template.HTML {
	return template.HTML(html)
}

func mbsubstr(text string, from, to int) string {

	rntext := []rune(text)

	if len(rntext) <= to {
		return text
	}

	return string(rntext[from:to])
}

func datestr(target time.Time) string {
	return target.Format("2006/01/02")
}
