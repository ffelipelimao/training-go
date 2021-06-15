package app

func Calculate(array []int32) []float64 {
	countInter := 0
	countNeg := 0
	countZero := 0
	var finalArray []float64

	lenght := len(array)

	for i := 0; i < lenght; i++ {
		if array[i] > 0 {
			countInter = countInter + 1
		}
		if array[i] < 0 {
			countNeg = countNeg + 1
		}
		if array[i] == 0 {
			countZero = countZero + 1
		}
	}

	finalInter := float64(countInter) / float64(lenght)
	finalNeg := float64(countNeg) / float64(lenght)
	finalZero := float64(countZero) / float64(lenght)

	finalArray = append(finalArray, finalInter, finalNeg, finalZero)

	return finalArray
}
