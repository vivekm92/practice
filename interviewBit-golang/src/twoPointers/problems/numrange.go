package twoPointers

/*
	Problem : https://www.interviewbit.com/problems/numrange/
*/

// T(n) : O(n) , S(n) : O(1)
func NumRange(A []int, B int, C int) int {

	return CountRange(A, C) - CountRange(A, B-1)
}

func CountRange(A []int, B int) int {

	count := 0
	start, end, n, sum := 0, 0, len(A), 0
	for end < n {
		sum += A[end]
		for start < n && sum > B {
			sum -= A[start]
			start++
		}
		end++
		count += (end - start) + 1
	}

	return count
}
