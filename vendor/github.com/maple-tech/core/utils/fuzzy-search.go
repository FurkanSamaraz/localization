package utils

import (
	"strings"
	"unicode/utf8"
)

//FuzzyDistance Uses Levenshtein distance between the query -> target string
//does so case-insensitive
//Returns the accuracy/distance as an int, if -1 then NO possible match, 0 is absolute match
func FuzzyDistance(query, target string) int {
	//Cleanup
	query = strings.ToLower(query)
	target = strings.ToLower(target)

	//Check lengths for easy exit
	lenDiff := len(target) - len(query)
	if lenDiff < 0 {
		return -1
	} else if lenDiff == 0 && query == target {
		return 0
	}

	//Levenshtein-ish algorithm
	charDiff := 0
Start:
	for _, cq := range query {
		for i, ct := range target {
			if cq == ct {
				target = target[i+utf8.RuneLen(ct):]
				continue Start
			} else {
				charDiff++
			}
		}
		return -1
	}

	return charDiff + utf8.RuneCountInString(target)
}

//FuzzyDistanceName Uses FuzzyDistance on a name by breaking up the spaces and re-arranging by number of parts
//If number of parts is 2 (ie. John Doe) then the name is reversed to Doe John
//If number of parts is more then 2 (ie. John Jackob Jingleheimer Smith) then the last part is first, followed by the beggining parts
//as in Smith John Jackob Jingleheimer
func FuzzyDistanceName(query, name string) int {
	parts := strings.Split(name, " ")
	if len(parts) > 1 {
		if len(parts) > 2 {
			end := len(parts) - 1
			return FuzzyDistance(query, parts[end]+" "+strings.Join(parts[:end-1], " "))
		}
		
		return FuzzyDistance(query, parts[1]+" "+parts[0])
	}
	return FuzzyDistance(query, name)
}
