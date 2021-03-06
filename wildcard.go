// Package `wildcard` provides utils for working with a classic wildcard matching (only `?` and `*` are supported).
//
// `Match` function is a golang's adaptation of Java's O(1) space solution:
//
// 		https://www.programcreek.com/2014/06/leetcode-wildcard-matching-java/
//
// and ~ 3.37 times faster then standard `filepath.Match`.
package wildcard

import "strings"

// Match - check if provided value match wildcard pattern
func Match(value, pattern string) bool {
	i := 0
	j := 0
	starIndex := -1
	iIndex := -1

	for i < len(value) {
		pLen := len(pattern)
		if j < pLen && (pattern[j] == '?' || pattern[j] == value[i]) {
			i++
			j++
		} else if j < pLen && pattern[j] == '*' {
			starIndex = j
			iIndex = i
			j++
		} else if starIndex != -1 {
			j = starIndex + 1
			i = iIndex + 1
			iIndex++
		} else {
			return false
		}
	}

	for j < len(pattern) && pattern[j] == '*' {
		j++
	}

	return j == len(pattern)
}

// IsPattern - check if provided value is a wildcard pattern,
// i.e. it has special characters "*" or "?"
func IsPattern(value string) bool {
	return strings.ContainsAny(value, "*?")
}
