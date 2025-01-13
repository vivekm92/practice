package arrayProblems

// T(n): O(n), S(n) : O(1)
func MaxArr(A []int) int {

	n, max1, min1, max2, min2 := len(A), -1<<63, 1<<63-1, -1<<63, 1<<63-1

	for i := 0; i < n; i++ {
		if max1 < A[i]+i {
			max1 = A[i] + i
		}

		if min1 > A[i]+i {
			min1 = A[i] + i
		}

		if max2 < A[i]-i {
			max2 = A[i] - i
		}

		if min2 > A[i]-i {
			min2 = A[i] - i
		}
	}

	res := 0
	if max1-min1 > max2-min2 {
		res = max1 - min1
	} else {
		res = max2 - min2
	}

	return res
}
