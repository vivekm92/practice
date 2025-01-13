package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/leap-year/
*/

// T(n) : O(1) , S(n) : O(1)
func LeapYear(A int) int {

	if A%4 != 0 {
		return 0
	}
	if A%100 == 0 && A%400 != 0 {
		return 0
	}
	return 1
}
