package mathPrimers

/*
	Problem : https://www.interviewbit.com/problems/mathbug02/
*/

// T(n) : O(n) , S(n) : O(1)
func SquareSum(A int) [][]int {

	ans := make([][]int, 0)
	for b := 1; b*b <= A; b++ {
		for a := 1; a <= b; a++ {
			if a*a+b*b == A {
				ans = append(ans, []int{a, b})
			}
		}
	}

	return ans
}
