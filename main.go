package main

import "fmt"

func main() {
	// queries := [][]int32{{8}, {1, 5}, {1, 6}, {3, 2}, {1, 10}, {1, 10}, {1, 6}, {2, 5}, {3, 2}}
	queries := [][]int32{{4}, {3, 4}, {2, 1003}, {1, 16}, {3, 1}}

	fmt.Println(freqQuery(queries))
}
