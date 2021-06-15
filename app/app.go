package app

func Calculate(array []int) int {
	result := 0
	for i := 0; i < len(array); i++ {
		result += i
	}
	return result
}
