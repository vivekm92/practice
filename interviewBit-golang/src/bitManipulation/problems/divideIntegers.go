package bitManipulation

/*
	Problem : https://www.interviewbit.com/problems/divide-integers/
*/

// T(n) : O(logn) , S(n) : O(1)
func DivideIntegers(A int, B int) int {

	if A == B {
		return 1
	}

	sign := 1
	if (A < 0 && B > 0) || (A > 0 && B < 0) {
		sign = -1
	}

	var res int
	a, b := abs(A), abs(B)
	for a >= b {
		count := 0
		for a > (b << (count + 1)) {
			count++
		}
		a -= (b << count)
		res += (1 << count)
	}

	if res == (1<<31) && sign == 1 {
		return (1 << 31) - 1
	}

	return res * sign
}

func abs(A int) int {
	if A >= 0 {
		return A
	}
	return -1 * A
}
