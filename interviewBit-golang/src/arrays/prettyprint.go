package arrays

/*
	problem : https://www.interviewbit.com/problems/prettyprint/
*/

// T(n) : O(n), S(n) : O(n)
func PrettyPrint(A int) [][]int {

	var sz int = 2*A - 1
	arr := make([][]int, sz)
	for i := 0; i < sz; i++ {
		arr[i] = make([]int, sz)
	}

	var mid int = sz / 2
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			d1, d2 := mid-i, mid-j
			if d1 < 0 {
				d1 = d1 * -1
			}
			if d2 < 0 {
				d2 = d2 * -1
			}
			m := d1
			if d2 > d1 {
				m = d2
			}
			arr[i][j] = m + 1
		}
	}

	return arr
}
