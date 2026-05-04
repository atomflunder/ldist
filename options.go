package ldist

import (
	"strings"
	"unicode"
)

// Option is a type alias for a function that takes two string pointers which should modify them in place.
type Option func(s1, s2 *string)

// ToLowercase converts both strings to lowercase.
func ToLowercase(s1, s2 *string) {
	*s1, *s2 = strings.ToLower(*s1), strings.ToLower(*s2)
}

// RemoveWhitespace removes all whitespace characters from both strings.
func RemoveWhitespace(s1, s2 *string) {
	removeWS := func(s string) string {
		var sb strings.Builder
		for _, r := range s {
			if !unicode.IsSpace(r) {
				sb.WriteRune(r)
			}
		}
		return sb.String()
	}
	*s1, *s2 = removeWS(*s1), removeWS(*s2)
}

// RemovePunctuation removes common punctuation characters from both strings.
func RemovePunctuation(s1, s2 *string) {
	punctuation := []string{".", ",", "!", "?", ";", ":", "-", "_", "(", ")", "[", "]", "{", "}", "\"", "'"}
	for _, p := range punctuation {
		*s1, *s2 = strings.ReplaceAll(*s1, p, ""), strings.ReplaceAll(*s2, p, "")
	}
}

// ToAlphanumeric removes all non-alphanumeric characters from both strings.
func ToAlphanumeric(s1, s2 *string) {
	var sb1, sb2 strings.Builder
	for _, r := range *s1 {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			sb1.WriteRune(r)
		}
	}
	for _, r := range *s2 {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			sb2.WriteRune(r)
		}
	}
	*s1, *s2 = sb1.String(), sb2.String()
}
