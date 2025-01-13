package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/kth-row-of-pascals-triangle/
*/

// T(n) : O(n) , S(n) : O(n)
func KthRowOfPascalTriangle(A int) []int {

	var prev, curr []int = []int{1}, []int{1}
	for A > 0 {
		n := len(prev)
		for i := 1; i < n; i++ {
			curr[i] = prev[i-1] + prev[i]
		}
		curr = append(curr, 1)
		prev = make([]int, 0)
		prev = append(prev, curr...)
		A--
	}

	return curr
}

// T(n) : O(n) , S(n) : O(n)
func kthRowOfPascalsTriangleRecursive(A int) []int {

	if A == 0 {
		return []int{1}
	}

	v, t := kthRowOfPascalsTriangleRecursive(A-1), make([]int, 0)
	t = append(t, 1)
	for i := 0; i < len(v)-1; i++ {
		t = append(t, v[i]+v[i+1])
	}
	t = append(t, 1)

	return t
}

// T(n) : O(n) ; S(n) : O(n)
func kthRowOfPascalsTriangleIterative(A int) []int {

	res := make([]int, 0)
	res = append(res, 1)
	if A == 0 {
		return res
	}

	for A > 0 {
		n, t := len(res), make([]int, 0)
		t = append(t, 1)
		for i := 0; i < n-1; i++ {
			t = append(t, res[i]+res[i+1])
		}
		t = append(t, 1)
		res = t
		A--
	}

	return res
}

// T(n) : O(n) ; S(n) : O(n)
func kthRowOfPascalsTriangle1(A int) []int {

	n := A
	row := make([]int, n+1)
	row[0], row[n] = 1, 1
	for i := 0; i < int(n/2); i++ {
		x := row[i] * (n - i) / (i + 1)
		row[i+1], row[n-i-1] = x, x
	}

	return row
}
