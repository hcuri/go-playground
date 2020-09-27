package main

import (
	"fmt"
)

func firstRepeatedWord(s string) string {

	myMap := make(map[string]int)
	delimiters := []rune{'\t', ' ', ',', ':', ';', '-', '.'}
	words := splitByDelimiter(s, delimiters)

	// create an empty map

	// split into array of words

	for _, val := range words {
		if _, ok := myMap[val]; ok {
			return val
		} else {
			myMap[val] = 1
		}
	}
	// iterate through words in s (make sure you are separting by delmiters)
	// check if word is in map, and return if yes
	// insert word to map

	fmt.Println(myMap)

	return ""
}

func splitByDelimiter(s string, delimiters []rune) []string {
	words := []string{}

	lastDelimLocation := 0

	//TODO make this more efficient with a map
	for i, char := range s {
		for _, del := range delimiters {
			if char == del {
				// this is a delimiter, store the word from lastDelimLocation until here
				// also update lasDelimLocation
				// if current location - last delim location > 1, then store the word
				if i-lastDelimLocation > 1 {
					words = append(words, s[lastDelimLocation:i])
				}
				lastDelimLocation = i + 1
			}
		}
	}
	// iterate over s

	// if rune is a delimiter, append the previous word to words array
	// use the [:] to store previous characters as a word

	return words
}
