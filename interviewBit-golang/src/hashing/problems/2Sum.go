package hashing

/*
	Problem : https://www.interviewbit.com/problems/2-sum/
*/

// T(n) : O(n), S(n) : O(n)
func TwoSum(A []int, B int) []int {

	n, lookup := len(A), make(map[int]int)
	for i := 0; i < n; i++ {
		if v, ok := lookup[B-A[i]]; ok {
			return []int{v + 1, i + 1}
		} else if _, found := lookup[A[i]]; !found {
			lookup[A[i]] = i
		}
	}

	return []int{}
}
