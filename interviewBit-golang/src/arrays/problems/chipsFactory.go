package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/chips-factory/
*/

// T(n) : O(n), S(n) : O(n)
func MoveZerosAtTheBack(A []int) []int {

	t, cnt := make([]int, 0), 0
	for _, v := range A {
		if v == 0 {
			cnt++
		} else {
			t = append(t, v)
		}
	}

	for cnt > 0 {
		t = append(t, 0)
		cnt--
	}

	return t
}
