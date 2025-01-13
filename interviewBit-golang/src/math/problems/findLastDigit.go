package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/find-last-digit/
*/

func findMod(A string, B int) int {
	mod := 0
	for _, v := range A {
		t := int(v - rune('0'))
		mod = (mod*10 + t) % B
	}
	return mod % B
}

// T(n) : O(n), S(n) : O(1)
func FindLastDigit(A string, B string) int {

	n, m := len(A), len(B)
	if n == 1 && m == 1 && A[0] == '0' && B[0] == '0' {
		return 1
	} else if n == 1 && A[0] == '0' {
		return 0
	} else if m == 1 && B[0] == '0' {
		return 1
	}

	exp := findMod(B, 4)
	if exp == 0 {
		exp = 4
	}

	res, p := 1, int(A[n-1]-'0')
	for exp > 0 {
		res = res * p
		exp--
	}

	return res % 10
}
