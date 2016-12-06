package string2

import (
	"bytes"
	"strings"
)

// Given a string, return a string where for every char in the original, there are two chars.
func DoubleChar(s string) string {
	var new []byte
	for i := range s {
		new = append(new, s[i], s[i])
	}
	return string(new)
}

// Return the number of times that the string "code" appears anywhere in the
// given string, except we'll accept any letter for the "d", so "cope" and "cooe" count.
// Try not to use regex
func CountCode(s string) int {
	if len(s) < 4 {
		return 0
	}
	prefix := []byte{'c', 'o'}
	suffix := []byte{'e'}
	count := 0
	for i := range s {
		if s[i] == prefix[0] && len(s[i:]) >= 4 {
			word := []byte(s[i : i+4])
			if bytes.HasPrefix(word, prefix) && bytes.HasSuffix(word, suffix) {
				count++
			}
		}
	}
	return count
}

// Return the number of times that the string "hi" appears anywhere in the given string.
func CountHi(s string) int {
	return strings.Count(s, "hi")
}

// Given two strings, return true if either of the strings appears at the very end
// of the other string, ignoring upper/lower case differences (in other words, the
//computation should not be "case sensitive"). Note: s.lower() returns the
// lowercase version of a string.
func EndOther(a, b string) bool {
	abytes := []byte(strings.ToLower(a))
	bbytes := []byte(strings.ToLower(b))
	if bytes.HasSuffix(abytes, bbytes) || bytes.HasSuffix(bbytes, abytes) {
		return true
	}
	return false
}

// Return true if the string "cat" and "dog" appear the same number of times
// in the given string.
func CatDog(s string) bool {
	if strings.Count(s, "cat") == strings.Count(s, "dog") {
		return true
	}
	return false
}

// Return true if the given string contains an appearance of "xyz" where the xyz
// is not directly preceeded by a period (.). So "xxyz" counts but "x.xyz" does not.
func XyzThere(s string) bool {
	if strings.Count(s, "xyz") > strings.Count(s, ".xyz") {
		return true
	}
	return false
}
