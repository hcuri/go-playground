package main

import (
	"strings"
)

func repeatedString(s string, n int64) int64 {
	betterN := int(n)

	// brute force solution
	// call function to get proper size string
	// count number of times a appears

	// better solution
	// count a's per repetition
	// figure put the a's in the last substring
	// multiply and add last one

	// calculate modulo of n / len(s)
	// that value is the size of the last string
	lenOfLastSubString := int(n) % len(s)

	// count a's in s till size of last string
	countOfLastSubstring := strings.Count(s[:lenOfLastSubString], "a")

	// count of base string
	countOfBaseString := strings.Count(s, "a")
	// result is n-1 * #as in string + #as in last string

	return int64(countOfBaseString*(betterN/len(s)) + countOfLastSubstring)
}
