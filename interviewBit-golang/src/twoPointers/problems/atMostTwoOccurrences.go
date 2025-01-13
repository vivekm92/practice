package twoPointers

/*
	Problem : https://www.interviewbit.com/problems/at-most-two-occurrences/
*/

func checkIfExists(A []int, m int) bool {

	n, st, ed := len(A), 0, 0
	lookup := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		ed = i
		k := ed - st + 1
		if k <= m {
			lookup[A[i]]++
			if lookup[A[i]] > 2 {
				return true
			}
		} else if k > m {
			lookup[A[i]]++
			if lookup[A[i]] > 2 && A[i] != A[st] {
				return true
			}
			if v, ok := lookup[A[st]]; ok {
				if v == 1 {
					delete(lookup, A[st])
				} else {
					lookup[A[st]]--
				}
			}
			st++
		}
	}

	return false
}

// T(n) : O(nlogn) , S(n) : O(n)
func AtMostTwoOccurrences(A []int) int {

	n := len(A)
	l, r := 0, n
	for l < r {
		m := l + (r-l)/2
		if checkIfExists(A, m) {
			r = m
		} else {
			l = m + 1
		}
	}

	return l - 1
}
