package greedy

/*
	Problem : https://www.interviewbit.com/problems/majority-element/
*/

// T(n) : O(n) ; S(n) : O(1)
func MajorityElement(A []int) int {

	var me, cnt int
	for _, v := range A {
		if cnt == 0 {
			me = v
			cnt++
		} else if me == v {
			cnt++
		} else {
			cnt--
		}
	}

	return me
}
