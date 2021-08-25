package main

import "fmt"

func main() {
	cores := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	result := breakingRecords(cores)
	fmt.Println(result)
}

func breakingRecords(scores []int32) []int32 {
	count := []int32{0, 0}
	high := scores[0]
	low := scores[0]

	for _, score := range scores {
		if score > high {
			high = score
			count[0]++
		} else if score < low {
			low = score
			count[1]++
		}
	}

	return count
}
