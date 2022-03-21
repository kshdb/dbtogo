package common

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func StrToCamelCase(src string) string {
	return strings.ReplaceAll(
		cases.Title(language.English).String(strings.ReplaceAll(src, "_", " ")),
		" ",
		"",
	)
}
