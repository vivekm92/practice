package stacksAndQueues

/*
	Problem : https://www.interviewbit.com/problems/maxspprod/
*/

// T(n) : O(n) , S(n) : O(n)
func Maxspprod(A []int) int {

	R := nextGreater(A)
	L := nextGreater(reverse(A))
	L = reverse(L)

	n := len(A)
	for i := 0; i < n; i++ {
		if L[i] != -1 {
			L[i] = n - 1 - L[i]
		}
	}

	var tRes int64
	for i := 0; i < n; i++ {
		if L[i] != -1 && R[i] != -1 {
			var curr int64 = int64(L[i]) * int64(R[i])
			if curr > tRes {
				tRes = curr
			}
		}
	}

	const MOD = 1000000007
	res := int(tRes % MOD)
	return res
}

func nextGreater(A []int) []int {

	n := len(A)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	st := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(st) == 0 || A[i] <= A[st[len(st)-1]] {
			st = append(st, i)
		} else {
			for len(st) > 0 && A[i] > A[st[len(st)-1]] {
				res[st[len(st)-1]] = i
				st = st[:len(st)-1]
			}
			st = append(st, i)
		}
	}

	return res
}

func reverse(A []int) []int {
	n := len(A)
	for i := 0; i < n/2; i++ {
		A[i], A[n-i-1] = A[n-i-1], A[i]
	}
	return A
}
