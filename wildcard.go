// Package wildcard provides the single function to check wildcard matching
package wildcard

// Match - checks if provided value match wildcard pattern
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
