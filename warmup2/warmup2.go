// Exercises from codingbat.com ported to Golang.
package warmup2

import (
	"bytes"
	"strings"
)

// Given an slice of ints, return true if the sequence of numbers 1, 2, 3
// appears in the slice somewhere.
func Array123(s []int) bool {
	l := len(s)
	if l < 3 {
		return false
	}
	for i, n := range s {
		if n == 1 && l-(i+1) >= 2 {
			if s[i+1] == 2 && s[i+2] == 3 {
				return true
			}
		}
	}
	return false
}

// Given a slice of ints, return the number of 9's in the slice.
func ArrayCount9(s []int) int {
	nines := 0
	for _, n := range s {
		if n == 9 {
			nines++
		}
	}
	return nines
}

// Given a slice of ints, return true if one of the first 4 elements in the
// slice is a 9. The slice length may be less than 4.
func ArrayFront(s []int) bool {
	var s4 []int
	if len(s) <= 4 {
		s4 = s
	} else {
		s4 = s[:4]
	}
	for _, n := range s4 {
		if n == 9 {
			return true
		}
	}
	return false
}

// Given a string and a non-negative int n, we'll say that the front of the
// string is the first 3 chars, or whatever is there if the string is less
// than length 3. Return n copies of the front
func FrontTimes(s string, n int) string {
	if len(s) <= 3 {
		return strings.Repeat(s, n)
	}
	return strings.Repeat(s[:3], n)
}

// Given a string, return the count of the number of times that a substring length
// 2 appears in the string and also as the last 2 chars of the string, so
// "hixxxhi" yields 1 (we won't count the end substring).
func Last2(s string) int {
	if len(s) <= 2 {
		return 0
	}
	last2 := s[len(s)-2:]
	count := 0
	for i, _ := range s {
		if i < len(s)-1 {
			if s[i] == last2[0] && s[i+1] == last2[1] {
				count++
			}
		}
	}
	// remove one match since the last2 don't count
	return count - 1
}

// Given a string, return a new string made of every other char starting with
// the first, so "Hello" yields "Hlo".
func StringBits(s string) string {
	var buffer bytes.Buffer
	for i, _ := range s {
		if i%2 == 0 {
			buffer.WriteByte(s[i])
		}
	}
	return buffer.String()
}

// Given 2 strings, a and b, return the number of the positions where they contain
// the same length 2 substring. So "xxcaazz" and "xxbaaz" yields 3, since the
// "xx", "aa", and "az" substrings appear in the same place in both strings.
func StringMatch(a, b string) int {
	var s1, s2 string
	var count int

	// find shortest so we iterate over it
	if len(a) <= len(b) {
		s1, s2 = a, b
	} else {
		s1, s2 = b, a
	}
	for i := 0; i < len(s1)-1; i++ {
		if s1[i:i+2] == s2[i:i+2] {
			count++
		}
	}
	return count
}

// Given a non-empty string like "Code" return a string like "CCoCodCode".
func StringSplosion(s string) string {
	var buffer bytes.Buffer
	for i := range s {
		buffer.WriteString(s[:i])
	}
	buffer.WriteString(s)
	return buffer.String()
}

// Given a string and a non-negative int n, return a larger string that is n
// copies of the original string.
func StringTimes(s string, n int) string {
	return strings.Repeat(s, n)
}
