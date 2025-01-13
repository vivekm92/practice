package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/multiple-left-rotations-of-the-array/
*/

func leftRotate(A []int, B int) []int {

	n, res := len(A), make([]int, 0)
	B = B % n
	for i := B; i < n; i++ {
		res = append(res, A[i])
	}

	for i := 0; i < B; i++ {
		res = append(res, A[i])
	}

	return res
}

// T(n) : O(B*A) ; S(n) : O(n)
func MultipleLeftRotations(A []int, B []int) [][]int {

	res := make([][]int, 0)
	for _, v := range B {
		arr := leftRotate(A, v)
		res = append(res, arr)
	}

	return res
}
