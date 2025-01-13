package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/sell-items/
*/

// T(n) : O(1) , S(n) : O(1)
func SellItems(A, B int) int {

	cnt := B / A
	if B%A != 0 {
		cnt++
	}

	nWeeks := cnt / 5
	if cnt%5 != 0 {
		nWeeks++
	}

	return nWeeks
}
