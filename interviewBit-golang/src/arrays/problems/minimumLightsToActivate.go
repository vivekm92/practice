package arrayProblems

// T(n) : O(n), S(n) : O(1)
func MinimumLightsToActivate(A []int, B int) int {

	curr, n, cnt := 0, len(A), 0
	for curr < n {
		prev := curr - B + 1
		nxt := curr + B - 1

		if prev < 0 {
			prev = 0
		}

		if nxt >= n {
			nxt = n - 1
		}

		idx := nxt
		for idx >= prev {
			if A[idx] == 1 {
				break
			}
			idx--
		}

		if idx < prev {
			return -1
		}

		cnt++
		curr = idx + B
	}

	return cnt
}
