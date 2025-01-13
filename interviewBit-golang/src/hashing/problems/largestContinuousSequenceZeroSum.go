package hashing

/*
	Problem : https://www.interviewbit.com/problems/largest-continuous-sequence-zero-sum/
*/

// Brute-Force Solution
// T(n) : O(n2) ; S(n) : O(1)
func LargestContinuousSequenceZeroSum1(A []int) []int {

	res, n, ml := make([]int, 0), len(A), 0
	for i := 0; i < n; i++ {
		t := 0
		for j := i; j < n; j++ {
			t += A[j]
			if t == 0 {
				l := j - i + 1
				if ml < l {
					res = A[i : j+1]
					ml = l
				}
			}
		}
	}

	return res
}

// Optimised Solution
// T(n) : O(n) ; S(n) : O(n)
func LargestContinuousSequenceZeroSum(A []int) []int {

	m, ml, t, res := make(map[int]int, 0), 0, 0, make([]int, 0)
	m[0] = -1
	for i, v := range A {
		t += v
		if _, ok := m[t]; ok {
			l := i - m[t]
			if ml < l {
				ml = l
				res = A[m[t]+1 : i+1]
			}
		} else {
			m[t] = i
		}
	}

	return res
}
