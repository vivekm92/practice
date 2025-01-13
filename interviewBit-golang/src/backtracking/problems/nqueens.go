package backtracking

/*
	Problem : https://www.interviewbit.com/problems/nqueens/
*/

func formatSolution(A int, B int) string {

	var res string
	for i := 0; i < A; i++ {
		if i == B {
			res += "Q"
		} else {
			res += "."
		}
	}

	return res
}

func NQueens(A int) [][]string {
	tempRes, B := make([][]int, 0), make([]int, 0)
	solveNQueens(A, 0, B, &tempRes)

	res := make([][]string, 0)
	for _, v1 := range tempRes {
		t := make([]string, 0)
		for _, v2 := range v1 {
			t = append(t, formatSolution(A, v2))
		}
		res = append(res, t)
	}
	return res
}

func solveNQueens(A int, row int, B []int, res *[][]int) {
	if row == A {
		t := make([]int, len(B))
		copy(t, B)
		*res = append(*res, t)
		return
	}

	for col := 0; col < A; col++ {
		B = append(B, col)
		if IsValidPlacement(B) {
			solveNQueens(A, row+1, B, res)
		}
		B = B[:len(B)-1]
	}
}

func IsValidPlacement(t []int) bool {

	rowId := len(t) - 1
	for i := 0; i < rowId; i++ {
		diff := abs(t[i] - t[rowId])
		if diff == 0 || diff == rowId-i {
			return false
		}
	}
	return true
}

func abs(A int) int {

	if A < 0 {
		return -1 * A
	}
	return A
}
