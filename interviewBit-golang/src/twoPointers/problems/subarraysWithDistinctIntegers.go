package twoPointers

/*
	Problem : https://www.interviewbit.com/problems/subarrays-with-distinct-integers/
*/

func solveSubarraysWithDistinctIntegers(A []int, B int) int {

	lookup := make(map[int]int, 0)
	j, totalCount, n := 0, 0, len(A)
	for i := 0; i < n; i++ {
		lookup[A[i]]++
		k := len(lookup)
		for j < n && k > B {
			if v, ok := lookup[A[j]]; ok {
				if v == 1 {
					delete(lookup, A[j])
				} else {
					lookup[A[j]]--
				}
			}
			j++
			k = len(lookup)
		}
		totalCount += (i - j + 1)
	}

	return totalCount
}

// T(n) : O(n) ; S(n) : O(n)
func SubarraysWithDistinctIntegers(A []int, B int) int {
	v1 := solveSubarraysWithDistinctIntegers(A, B)
	v2 := solveSubarraysWithDistinctIntegers(A, B-1)

	return v1 - v2
}
