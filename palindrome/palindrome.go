package palindrome

import (
	"regexp"
	"strings"
)

var nonAlphaNumeric = regexp.MustCompile("[^a-zA-Z0-9]+")

func Is(phrase string) bool {
	// Remove all alphanumerics and convert to lowercase.
	phrase = strings.ToLower(
		nonAlphaNumeric.ReplaceAllString(phrase, ""))

	for i := 0; i < len(phrase)/2; i++ {
		if phrase[i] != phrase[len(phrase)-i-1] {
			return false
		}
	}
	return true
}
