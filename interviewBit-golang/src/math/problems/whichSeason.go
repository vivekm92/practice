package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/which-season/
*/

// T(n) : O(1) , S(n) : O(1)
func WhichSeason(A int) string {

	if A < 1 || A > 12 {
		return "Invalid"
	} else if A <= 2 || A >= 12 {
		return "Winter"
	} else if A >= 9 && A <= 11 {
		return "Autumn"
	} else if A >= 6 && A <= 8 {
		return "Summer"
	} else {
		return "Spring"
	}
}
