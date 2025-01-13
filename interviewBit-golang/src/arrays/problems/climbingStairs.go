package arrayProblems

/*
	problem : https://www.interviewbit.com/problems/climbing-stairs/
*/

// T(n) : O(n), S(n) : O(n)
func ClimbingStairs(A []int) int {

	n := len(A)
	t := make([]int, n+1)
	t[0], t[1] = A[0], A[0]
	for i := 2; i <= n; i++ {
		c1, c2 := A[i-1]+t[i-1], A[i-1]+t[i-2]
		t[i] = c1
		if c2 < c1 {
			t[i] = c2
		}
	}

	return t[n]
}
