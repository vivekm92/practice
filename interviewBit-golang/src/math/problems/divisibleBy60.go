package mathProblems

func DivisibleBy60(A []int) int {

	var zeroPresent bool
	var evenCount int
	var numSum int
	for _, v := range A {
		if v == 0 {
			zeroPresent = true
		}
		if v%2 == 0 {
			evenCount++
		}
		numSum += v
	}

	if len(A) == 1 && A[0] == 0 {
		return 1
	}

	if zeroPresent && evenCount > 1 && numSum%3 == 0 {
		return 1
	}

	return 0
}
