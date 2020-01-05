package main

// 1.1

// HasAllUniqueCharacters checks if a string is comprised of only unique characters
func HasAllUniqueCharacters(s string) bool {
	counts := make(map[rune]int)

	for _, c := range s {
		counts[c] = counts[c] + 1
	}

	hasDuplicate := false

	for _, count := range counts {
		if count > 1 {
			hasDuplicate = true
		}
	}

	return !hasDuplicate
}

// 1.2

// Reverse returns a string with the characters of the input string in reverse order
func Reverse(s string) string {
	length := len(s)
	buff := make([]rune, length)

	for i, r := range s {
		buff[length-1-i] = r
	}

	var reversed string

	for _, r := range buff {
		reversed = reversed + string(r)
	}

	return reversed
}
