package timeComplexity

/*
	Problem : https://www.interviewbit.com/problems/collatz-conjecture/
*/

// T(n) : O(n), S(n) : O(1)
func Collatz(A int, B int) int64 {

	res := int64(A)
	for B > 1 {
		if res%2 == 0 {
			res /= 2
		} else {
			res = 3*res + 1
		}
		B--
	}

	return res
}
