package main

import "strings"

func stringMatching(words []string) []string {
	substrings := []string{}
	wordSet := make(map[string]bool)
	substringSet := make(map[string]bool)

	// add all words to wordSet
	for _, word := range words {
		// if word is not in wordSet add it
		if !wordSet[word] {
			wordSet[word] = true
		}
	}

	// iterate over the words
	for _, word := range words {
		// if word is in substringSet
		if substringSet[word] {
			// continue
			continue
		}

		// iterate over wordsSet
		for existingWord := range wordSet {
			// if word is the same continue
			if word == existingWord {
				continue
			}

			// if word is a substring of current word
			if isSubstring(existingWord, word) {
				// add word to substringSet and substrings
				substringSet[word] = true
			}
		}

	}

	for word := range substringSet {
		substrings = append(substrings, word)
	}

	return substrings
}

func isSubstring(word1 string, word2 string) bool {
	// iterate over words in
	return strings.Contains(word1, word2)
}
