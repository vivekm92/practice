package arrayProblems

/*
  Problem : https://www.interviewbit.com/problems/min-steps-in-infinite-grid/

  Solution :
    - All the points need to be covered in order ( in which they are given ).
	- since we can move in diagonal as well in cartesian plane, when we move from (x,y) to (a,b)
	- minimum steps to move from (x,y) --> (a,b) would be max(abs(a-x), abs(b-y))
*/

// T(n) : O(n) , S(n) : O(1)
func minStepsInInfiniteGrid(A []int, B []int) int {

	n, res := len(A), 0
	for i := 0; i < n-1; i++ {
		x, y := abs(A[i]-A[i+1]), abs(B[i]-B[i+1])
		if x > y {
			res += x
		} else {
			res += y
		}
	}

	return res
}

func abs(A int) int {
	if A < 0 {
		return -1 * A
	}
	return A
}
