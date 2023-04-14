package utils

import (
	"bytes"
	"html"
	"html/template"
	"strings"
)

//SanitizeHTML escapes and removes any HTML in the string
//returning the clean version.
func SanitizeHTML(str string) string {
	var out string
	if !strings.ContainsAny(str, "<>") {
		out = str //Already clean
	} else {
		var buf bytes.Buffer
		inTag := false
		var tag string
		for _, r := range str {
			switch r {
			case '<':
				inTag = true
				tag = ""
			case '>':
				inTag = false
				//Check what the tag was
				tag = strings.ToLower(tag)
				if tag == "p" || tag == "br" {
					buf.WriteRune('\n')
				}
			case ' ', '\n', '/':
				if !inTag {
					buf.WriteRune(r)
				}
			default:
				if inTag {
					tag = tag + string(r)
				} else {
					buf.WriteRune(r)
				}
			}
		}
		out = buf.String()
	}

	out = html.UnescapeString(out)
	out = template.HTMLEscapeString(out)
	return out
}
