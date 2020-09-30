package main

import (
	"fmt"
)

func main() {

	// add input here

	// sentence := "We ;work .hard \tbecause hard work pays."

	words := []string{"leetcoder", "leetcode", "od", "hamlet", "am"}

	// result := splitByDelimiter(sentence, delimiters)

	result := stringMatching(words) // insert function call here

	fmt.Println(result)

}
