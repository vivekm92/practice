package bsearch

/*
	Problem : https://www.interviewbit.com/problems/implement-power-function/
*/

func solvePow(a int64, b int64, c int64) int64 {

	if b == 0 {
		return 1
	} else if b == 1 {
		return a % c
	}

	if b%2 != 0 {
		return ((a % c) * (solvePow(a, b-1, c) % c)) % c
	}

	k := solvePow(a, b/2, c)
	return ((k % c) * (k % c)) % c
}

// T(n) : O(logn), S(n) : O(1)
func Pow(A int, B int, C int) int {
	a := int64(A)
	b := int64(B)
	c := int64(C)

	a = a % c
	if a == 0 {
		return 0
	}

	var r int64 = solvePow(a, b, c)
	if r < 0 {
		return int(r + c)
	}

	return int(r)
}
