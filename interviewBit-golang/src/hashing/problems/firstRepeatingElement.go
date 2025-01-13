package hashing

// Brute-Force Solution
// T(n) : O(n2) ; S(n) : O(1)
func FirstRepeatingElement1(A []int) int {

	res, n := -1, len(A)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if A[j] == A[i] {
				return A[i]
			}
		}
	}

	return res
}

// Optimised Solution
// T(n) : O(n) ; S(n) : O(n)
func FirstRepeatingElement(A []int) int {

	m, tidx, res := make(map[int]int, 0), len(A), -1
	for i, v := range A {
		if idx, ok := m[v]; ok {
			if idx < tidx {
				tidx = idx
				res = v
			}
		} else {
			m[v] = i
		}
	}

	return res
}
