package arrayProblems

// T(n) : O(n), S(n) : O(1)
func MoveZeroes(A []int) []int {

	iA, iB := 0, 0
	for iB < len(A) {
		if A[iB] != 0 {
			A[iA], A[iB] = A[iB], A[iA]
			iA++
		}
		iB++
	}

	return A
}
