/* Problem Statement
Implement a profanity filter system.
Filter(s) that takes a string and filter it using a word list
Assume the list is available in plain text file.

*/

/* Edge Cases and assumptions
Does casing mater? Aggro vs aggro
*/

/* Potential Solutions
1. Create a hashset of the badWords
*/

package main

import (
	"strings"
	//"io/ioutil"
)

// func main (){
//     sentence := "filter aggro this aggro word"
//     badWords := map[string]bool{}

//     // TODO read in file into forbiddenWords
//     // TEMP simulate adding things to the map
//     badWords["aggro"] = true

//     // file io
//     // data, err := ioutil.ReadFile()
//     // check(err)

//     // sentence = string(data)

//     fmt.Println(filter(sentence, badWords))
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func filter(sentence string, badWords map[string]bool) string {
	filtered := ""

	// split string into array of strings
	sentenceArray := strings.Split(sentence, " ")

	// iterate over the sentence and store only good words
	for _, word := range sentenceArray {
		if _, ok := badWords[word]; ok {
			continue
		}
		filtered = filtered + " " + word
	}

	// join resulting string

	return filtered
}
