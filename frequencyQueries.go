package main

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {

	var type3outputs []int32

	// create map to store values and repetitions
	var m = make(map[int32]int32)

	// iterate through commands
	for _, query := range queries[1:] {
		switch query[0] {
		case 1:
			// increment count of this value
			m[query[1]] = m[query[1]] + 1
		case 2:
			// decrement count of this value, unless it's already zero
			if m[query[1]] > 0 {
				m[query[1]] = m[query[1]] - 1
			}
		case 3:
			insert := false
			// iterate through map
			for i := range m {
				// if any value is equal to query[1]
				if m[i] == query[1] {
					insert = true
					break
				}
			}
			if insert {
				type3outputs = append(type3outputs, 1)
			} else {
				type3outputs = append(type3outputs, 0)
			}
		}
	}

	// fmt.Println(m)
	// if 1

	return type3outputs
}
