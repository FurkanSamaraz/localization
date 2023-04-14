package utils

import "strings"

//ParseBool takes a string and returns true
//if the first rune is 1 or t (case-insensitive)
func ParseBool(str string) bool {
	if len(str) == 0 {
		return false
	}

	c := []rune(strings.ToLower(str))[0]
	if c == '1' || c == 't' {
		return true
	}
	return false
}
