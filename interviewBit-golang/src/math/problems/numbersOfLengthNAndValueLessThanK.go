package mathProblems

import (
	"math"
)

func digitArray(C int) []int {

	var arr []int
	for C != 0 {
		arr = append(arr, C%10)
		C = C / 10
	}

	n := len(arr)
	for i := 0; i < n/2; i++ {
		t := arr[n-1-i]
		arr[n-1-i] = arr[i]
		arr[i] = t
	}

	return arr
}

func NumbersOfLengthNAndValueLessThanK(A []int, B int, C int) int {

	n := len(A)
	digit := digitArray(C)
	// Case 1
	if n == 0 || len(digit) < B {
		return 0
	}

	// Case 2
	if len(digit) > B {
		if A[0] == 0 && B != 1 {
			return int((n - 1) * int(math.Pow(float64(n), float64(B-1))))
		} else {
			return int(math.Pow(float64(n), float64(B)))
		}
	}

	// Case 3
	lower := make([]int, 11)
	for i := 0; i < n; i++ {
		lower[A[i]+1] = 1
	}

	for i := 1; i < 11; i++ {
		lower[i] += lower[i-1]
	}

	dp := make([]int, B+1)
	var flag bool = true
	dp[0] = 0
	for i := 1; i <= B; i++ {
		d2 := lower[digit[i-1]]
		dp[i] = dp[i-1] * n

		if i == 1 && A[0] == 0 && B != 1 {
			d2 -= 1
		}

		if flag {
			dp[i] += d2
		}

		flag = flag && (lower[digit[i-1]+1] == (lower[digit[i-1]] + 1))
	}

	return dp[B]
}
