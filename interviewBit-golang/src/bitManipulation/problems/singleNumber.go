package bitManipulation

/*
problem : https://www.interviewbit.com/problems/single-number/
*/

// T(n) : O(n), S(n) : O(1)
func SingleNumber(A []int) int {

	var res int
	for _, v := range A {
		res = res ^ v
	}

	return res
}
