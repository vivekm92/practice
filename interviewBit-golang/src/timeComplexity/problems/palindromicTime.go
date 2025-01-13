package timeComplexity

/*
	Problem : https://www.interviewbit.com/problems/palindromic-time/
*/

// T(n) : O(1) , S(n) : O(1)
func PalindromicTime(A string) int {

	hour := int(A[0]-'0')*10 + int(A[1]-'0')
	minute := int(A[3]-'0')*10 + int(A[4]-'0')
	count := 0
	for (hour/10 != minute%10) || (hour%10 != minute/10) {
		minute++
		count++
		if minute == 60 {
			minute = 0
			hour++
		}
		if hour == 24 {
			hour = 0
		}
	}

	return count
}
