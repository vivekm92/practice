package arrayProblems

/*
	problem : https://www.interviewbit.com/problems/flip/
*/

// T(n) : O(n), S(n) : O(n)
func Flip(A string) []int {

	n, arr, flag := len(A), make([]int, 0), true
	for i := 0; i < n; i++ {
		if A[i] == '0' {
			arr = append(arr, 1)
			flag = false
		} else {
			arr = append(arr, -1)
		}
	}

	if flag {
		return []int{}
	}

	currSum, maxSum, l, res := 0, 0, 0, make([]int, 2)
	for i := 0; i < n; i++ {
		currSum += arr[i]
		if currSum > maxSum {
			maxSum = currSum
			res[0] = l + 1
			res[1] = i + 1
		}

		if currSum < 0 {
			currSum = 0
			l = i + 1
		}
	}

	return res
}
