package arrays

/*
	problem : https://www.interviewbit.com/problems/arraybug/
*/

// T(n) : O(n), S(n) : O(n)
func RotateArray(A []int, B int) []int {

	n := len(A)
	ret := make([]int, 0)

	for i := 0; i < n; i++ {
		ret = append(ret, A[(i+B)%n])
	}

	return ret
}
