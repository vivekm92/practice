package arrayProblems

// T(n) : O (n), S(n) : O(n)
func PlusOne(A []int) []int {

	carry := 1
	for i := len(A) - 1; i >= 0; i-- {

		A[i] += carry
		carry = A[i] / 10
		A[i] = A[i] % 10

		if carry == 0 {
			break
		}
	}

	var res = []int{}
	if carry != 0 {
		res = append(res, carry)
		for i := 0; i < len(A); i++ {
			res = append(res, A[i])
		}
	} else {
		i := 0
		for A[i] == 0 && i < len(A) {
			i++
		}
		for i < len(A) {
			res = append(res, A[i])
			i++
		}

	}

	return res
}
