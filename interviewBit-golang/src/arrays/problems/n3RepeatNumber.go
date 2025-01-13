package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/n3-repeat-number/
*/

// T(n) : O(n), S(n) : O(1)
func RepeatNumber(A []int) int {

	count1, count2 := 0, 0
	num1, num2 := -1, -1
	for _, a := range A {
		if a == num1 {
			count1++
		} else if a == num2 {
			count2++
		} else if count1 == 0 {
			num1 = a
			count1++
		} else if count2 == 0 {
			num2 = a
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, a := range A {
		if a == num1 {
			count1++
		} else if a == num2 {
			count2++
		}
	}

	n := len(A)
	if count1 > n/3 {
		return num1
	} else if count2 > n/3 {
		return num2
	}
	return -1
}
