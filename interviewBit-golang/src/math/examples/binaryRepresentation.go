package mathExamples

// T(n): O(n), S(n) : O(1), where n = log(A)
func FindDigitsInBinary(A int) string {

	if A == 0 {
		return "0"
	}

	k, n := 1, 0
	for k <= A {
		k = k * 2
		n += 1
	}

	k = k / 2
	var res string
	for i := 0; i < n; i++ {
		if A >= k {
			A = A - k
			res += "1"
		} else {
			res += "0"
		}
		k = k / 2
	}

	return res
}
