package utils

import (
	"bytes"
	"regexp"
	"unicode"
)

var titleRegex = regexp.MustCompile(`([^\s]+)`)

//TitleCaseString turns a proper-name/title into it's english/latin based capitalization
//Example: " sOme title   sTRIng " -> "Some Title String"
//TODO: Handle locale name capitalization ie. "McMillian"
func TitleCaseString(str string) string {
	parts := titleRegex.FindAllString(str, -1)
	var b bytes.Buffer
	for i, part := range parts {
		for i, char := range part {
			if i == 0 {
				b.WriteRune(unicode.ToUpper(char))
			} else {
				b.WriteRune(unicode.ToLower(char))
			}
		}

		if i < len(parts)-1 {
			b.WriteRune(' ')
		}
	}

	return b.String()
}
