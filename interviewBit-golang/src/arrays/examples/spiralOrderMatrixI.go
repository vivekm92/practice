package arrayExamples

/*
	Problem : https://www.interviewbit.com/problems/spiral-order-matrix-i/
	LeetCode : https://leetcode.com/problems/spiral-matrix/
*/

// T(n) : O(n); S(n) : O(n)
func SpiralOrder1(A [][]int) []int {

	// rs, re : row start, end
	// cs, ce : col start, end
	rs, re := 0, len(A)-1
	cs, ce := 0, len(A[0])-1

	res := make([]int, 0)
	for rs <= re || cs <= ce {
		// top row
		for i := cs; i <= ce; i++ {
			res = append(res, A[rs][i])
		}
		rs++
		if rs > re {
			return res
		}

		// right col
		for i := rs; i <= re; i++ {
			res = append(res, A[i][ce])
		}
		ce--
		if cs > ce {
			return res
		}

		// bottom row
		for i := ce; i >= cs; i-- {
			res = append(res, A[re][i])
		}
		re--
		if rs > re {
			return res
		}

		// left col
		for i := re; i >= rs; i-- {
			res = append(res, A[i][cs])
		}
		cs++
		if cs > ce {
			return res
		}

	}

	return res
}
