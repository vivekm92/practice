package arrayProblems

/*
  Problem : https://www.interviewbit.com/problems/pick-from-both-sides/
  Similar Problem : https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/

  Solution :
    - sum first k elements of array
	- find max sum that can be achieved by removing 1 element and adding 1 element from each ends.

*/

// T(n) : O(n) , S(n) : O(1)
func pickFromBothSides(A []int, B int) int {

	n, t := len(A), 0
	for i := 0; i < B; i++ {
		t += A[i]
	}

	res := t
	for i := 0; i < B; i++ {
		t += A[n-1-i] - A[B-1-i]
		if t > res {
			res = t
		}
	}

	return res
}
