package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/string-inversion/
*/

// T(n) : O(n) , S(n) : O(n)
func StringInversion(A string) string {

	res := ""
	for _, v := range A {
		if v >= 'a' && v <= 'z' {
			res += string(v - 32)
		} else {
			res += string(v + 32)
		}
	}

	return res
}
