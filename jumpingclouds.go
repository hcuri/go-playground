package main

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {

	jumps := 0
	position := 0

	for {
		// if position == len(c) - 1, return jumps
		// elif position == len(c) - 2, position++, jumps++
		// elif position + 2 == 0, position+=2, jumps++
		// else position++, jumps++

		if position == len(c)-1 {
			return int32(jumps)
		} else if position == len(c)-2 {
			position++
			jumps++
		} else if c[position+2] == 0 {
			position += 2
			jumps++
		} else {
			position++
			jumps++
		}
	}
}
