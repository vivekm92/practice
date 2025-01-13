package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/odd-even-rule/
*/

// T(n) : O(n) , S(n) : O(1)
func OddEvenRule(A []int, B int, C int) int {

	oddDay := false
	if B%2 != 0 {
		oddDay = true
	}

	fineCollected := 0
	for _, v := range A {
		if (oddDay && v%2 == 0) || (!oddDay && v%2 != 0) {
			fineCollected += C
		}
	}

	return fineCollected
}
