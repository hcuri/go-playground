package main

import (
	"fmt"
)

func main() {

	// add input here

	sentence := "We ;work .hard \tbecause hard work pays."

	// result := splitByDelimiter(sentence, delimiters)

	result := firstRepeatedWord(sentence) // insert function call here

	fmt.Println(result)

}
