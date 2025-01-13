package twoPointers

/*
	problem : https://www.interviewbit.com/problems/remove-element-from-array/
*/

// T(n) : O(n), S(n) : O(1)
func RemoveElements(A *[]int, B int) int {

	cnt := 0
	for _, v := range *A {
		if v != B {
			(*A)[cnt] = v
			cnt++
		}
	}

	return cnt
}
