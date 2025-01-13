package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/sum-of-7-s-multiple/
*/

// T(n) : O(1) , S(n) : O(1)
func SumOf7sMultiple(A int, B int) int64 {

	a := A
	if A%7 != 0 {
		a += (7 - A%7)
	}

	b := B
	if B%7 != 0 {
		b -= B % 7
	}

	n := (b-a)/7 + 1
	return int64(n) * int64(a+b) / 2
}
