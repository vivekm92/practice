package arrayProblems

// T(n) : O(n), S(n) : O(n)
func CountPartitions(A int, B []int) int {

	n, sum := len(B), 0
	for i := 0; i < n; i++ {
		sum += B[i]
	}

	if sum%3 != 0 {
		return 0
	}

	target, ts := sum/3, 0
	cnt := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		ts += B[i]
		if ts == target {
			cnt[i] = 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		cnt[i] += cnt[i+1]
	}

	ts = 0
	res := 0
	for i := 0; i < n-2; i++ {
		ts += B[i]
		if ts == target {
			res += cnt[i+2]
		}
	}

	return res
}
