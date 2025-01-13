package greedy

/*
	Problem : https://www.interviewbit.com/problems/gas-station/
*/

// T(n) : O(n) , S(n) : O(1)
func GasStation(A []int, B []int) int {

	n, totalGas, totalCost := len(A), 0, 0
	for i := 0; i < n; i++ {
		totalGas += A[i]
		totalCost += B[i]
	}

	if totalGas < totalCost {
		return -1
	}

	res, currGas, currCost := 0, 0, 0
	for i := 0; i < n; i++ {
		currGas += A[i]
		currCost += B[i]

		if currGas < currCost {
			res = i + 1
			currGas = 0
			currCost = 0
		}
	}

	return res
}
