package utils

import (
	"net/http"

	"golang.org/x/text/language"
)

var matcher = language.NewMatcher([]language.Tag{
	language.English,
	language.Turkish,
})

//ParseLanguage takes an incoming language and matches it against the ACCEPTED list of languages
//Return is only the BSP Base, ie: en-Us -> en
func ParseLanguage(lang string) string {
	tag, _ := language.MatchStrings(matcher, lang)
	base, _ := tag.Base()
	return base.String()
}

//ParseLanguageHTTPFallback takes an incoming language, AND the request and matches it against the
//ACCEPTED list of language. If the initial string doesn't match, the http.Request is used as a fallback
//for it's Accept-Language header. Return is only the BSP Base, ie: en-Us -> en
func ParseLanguageHTTPFallback(lang string, r *http.Request) string {
	accept := r.Header.Get("Accept-Language")
	tag, _ := language.MatchStrings(matcher, lang, accept)
	base, _ := tag.Base()
	return base.String()
}
