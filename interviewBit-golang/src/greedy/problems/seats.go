package greedy

/*
	Problem : https://www.interviewbit.com/problems/seats/
*/

// T(n) : O(n) ; S(n) : O(n)
func Seats(A string) int {

	arr := make([]int, 0)
	for i, v := range A {
		if v == 'x' {
			arr = append(arr, i)
		}
	}

	for i, v := range arr {
		arr[i] = v - i
	}

	n := len(arr)
	if n == 0 {
		return 0
	}

	const MOD = 10000003
	ss, total := arr[n/2], 0
	for _, v := range arr {
		diff := ss - v
		if diff < 0 {
			diff = -1 * diff
		}
		total = (total%MOD + diff%MOD) % MOD
	}

	return total
}
