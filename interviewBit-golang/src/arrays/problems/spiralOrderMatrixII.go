package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/spiral-order-matrix-ii/
*/

// T(n) : O(n2); S(n) : O(n2)
func GenerateMatrix(A int) [][]int {

	m := make([][]int, A)
	for i := 0; i < A; i++ {
		m[i] = make([]int, A)
	}

	t := 1
	rs, re := 0, A-1
	cs, ce := 0, A-1
	for t <= A*A {

		for i := cs; i <= ce; i++ {
			m[rs][i] = t
			t++
		}
		rs++
		if t > A*A || rs > re {
			return m
		}

		for i := rs; i <= re; i++ {
			m[i][ce] = t
			t++
		}
		ce--
		if t > A*A || cs > ce {
			return m
		}

		for i := ce; i >= cs; i-- {
			m[re][i] = t
			t++
		}
		re--
		if t > A*A || rs > re {
			return m
		}

		for i := re; i >= rs; i-- {
			m[i][cs] = t
			t++
		}
		cs++
		if t > A*A || cs > ce {
			return m
		}
	}

	return m
}
