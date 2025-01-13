package bitManipulation

/*
	Problem : https://www.interviewbit.com/problems/single-number-ii/
*/

// T(n) : O(n) , S(n) : O(1)
func SingleNumberII(A []int) int {

	res := 0
	for i := 0; i < 32; i++ {
		sum := 0
		for _, v := range A {
			sum += (v >> i) & 1
		}
		sum %= 3
		res |= (sum << i)
	}

	return res
}
