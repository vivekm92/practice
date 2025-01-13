package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/number-of-sundays/
*/

// T(n) : O(1) , S(n) : O(1)
func NumberOfSundays(A string, B int) int {

	x := -1
	var lookup []string = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i, v := range lookup {
		if A == v {
			x = i + 1
			break
		}
	}

	y := 7 - x
	if y >= B {
		return 0
	}

	return ((B - y - 1) / 7) + 1
}
