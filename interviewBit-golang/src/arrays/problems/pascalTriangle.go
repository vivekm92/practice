package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/pascal-triangle/
*/

func solvePascalTriangle(res [][]int, A int) [][]int {
	if A == 1 {
		res[0] = []int{1}
		return res
	} else if A == 2 {
		res[0] = []int{1}
		res[1] = []int{1, 1}
		return res
	}

	res = solvePascalTriangle(res, A-1)
	arr := make([]int, A)
	arr[0] = 1
	for i := 1; i < A; i++ {
		arr[i] = res[A-2][i-1]
		if i < A-1 {
			arr[i] += res[A-2][i]
		}
	}
	res[A-1] = arr
	return res
}

// T(n) : O(n*m), S(n) : (n*m)
func PascalTriangle(A int) [][]int {
	if A == 0 {
		return [][]int{}
	}
	res := make([][]int, A)
	return solvePascalTriangle(res, A)
}
