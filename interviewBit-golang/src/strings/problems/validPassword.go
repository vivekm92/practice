package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/valid-password/
*/

// T(n) : O(n), S(n) : O(1)
func IsValidPassword(A string) int {

	var containsNumerical, containsLowerCase, containsUpperCase, containsSpecialCharacter bool
	for _, v := range A {
		if v >= '0' && v <= '9' {
			containsNumerical = true
		} else if v >= 'a' && v <= 'z' {
			containsLowerCase = true
		} else if v >= 'A' && v <= 'Z' {
			containsUpperCase = true
		} else if v == '@' || v == '#' || v == '%' || v == '&' || v == '!' || v == '$' || v == '*' {
			containsSpecialCharacter = true
		}
	}

	flag := containsNumerical && containsLowerCase && containsUpperCase &&
		containsSpecialCharacter && len(A) >= 8 && len(A) <= 15

	if flag {
		return 1
	}

	return 0
}
