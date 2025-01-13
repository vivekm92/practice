package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/bribing-was-never-easy/
*/

// T(n) : O(n) , S(n) : O(1)
func BribingWasNeverEasy(A []int, B []int) int {

	totalBribes, n := 0, len(A)
	for i := n - 1; i >= 0; i-- {
		currPos, actualPos := i+1, A[i]
		// if currPos < actualPos {
		diff := currPos - actualPos
		if B[actualPos-1] < diff {
			return -1
		} else {
			var t int = A[currPos-1]
			A[currPos-1] = A[actualPos-1]
			var a []int = append(append(A[:actualPos-1], t), A[actualPos:]...)
			A = a
			totalBribes += diff
		}
		// }
	}

	return totalBribes
}
