package greedy

// Problem : https://www.interviewbit.com/problems/interview-questions/

// T(n) : O(n) ; S(n) : O(1)
func Bulbs(A []int) int {

	cnt := 0
	for _, v := range A {
		if (v == 0 && cnt%2 == 0) || (v == 1 && cnt%2 == 1) {
			cnt++
		}
	}

	return cnt
}
