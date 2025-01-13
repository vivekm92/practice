package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/multiply-strings/
*/

// T(n) : O(n2) , S(n) : O(1)
func MultiplyStrings(A string, B string) string {

	if A == "0" || B == "0" {
		return "0"
	}

	A = removeZeros(A)
	B = removeZeros(B)
	a1 := reverse([]rune(A))
	a2 := reverse([]rune(B))

	n := len(a1) + len(a2)
	res := make([]rune, n)
	for i := range res {
		res[i] = '0'
	}

	for p2 := range a2 {
		d2 := a2[p2] - '0'
		for p1 := range a1 {
			d1 := a1[p1] - '0'
			pos := p1 + p2
			carry := res[pos] - '0'
			m := d1*d2 + carry
			res[pos] = rune(m%10) + '0'
			res[pos+1] += rune(m / 10)
		}
	}

	res = reverse(res)
	ans := string(res)
	ans = removeZeros(ans)
	return ans
}

func reverse(A []rune) []rune {
	n := len(A)
	for i := 0; i < n/2; i++ {
		A[i], A[n-i-1] = A[n-i-1], A[i]
	}
	return A
}

func removeZeros(A string) string {
	i := 0
	for i < len(A) && A[i] == '0' {
		i++
	}
	return A[i:]
}
